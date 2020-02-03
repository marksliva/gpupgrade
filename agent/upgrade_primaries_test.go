package agent

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/testutils/exectest"
	"github.com/greenplum-db/gpupgrade/utils"
)

// Does nothing.
func Success() {}

// Prints the environment, one variable per line, in NAME=VALUE format.
func EnvironmentMain() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

func FailedMain() {
	os.Exit(1)
}

func init() {
	exectest.RegisterMains(
		Success,
		EnvironmentMain,
		FailedMain,
	)
}

func TestUpgradePrimary(t *testing.T) {
	// Disable exec.Command. This way, if a test forgets to mock it out, we
	// crash the test instead of executing code on a dev system.
	execCommand = nil

	// We need a real temporary directory to change to. Replace MkdirAll() so
	// that we can make sure the directory is the correct one.
	tempDir, err := ioutil.TempDir("", "gpupgrade")
	if err != nil {
		t.Fatalf("creating temporary directory: %+v", err)
	}
	defer os.RemoveAll(tempDir)

	utils.System.MkdirAll = func(path string, perms os.FileMode) error {
		// Bail out if the implementation tries to touch any other directories.
		if !strings.HasPrefix(path, tempDir) {
			t.Fatalf("requested directory %q is not under temporary directory %q; refusing to create it",
				path, tempDir)
		}

		return os.MkdirAll(path, perms)
	}
	defer func() {
		utils.System = utils.InitializeSystemFunctions()
	}()

	pairs := []*idl.DataDirPair{
		{
			SourceDataDir: "/data/old",
			TargetDataDir: "/data/new",
			SourcePort:    15432,
			TargetPort:    15433,
			Content:       1,
			DBID:          2,
		},
		{
			SourceDataDir: "/other/data/old",
			TargetDataDir: "/other/data/new",
			SourcePort:    99999,
			TargetPort:    88888,
			Content:       7,
			DBID:          6,
		},
	}

	// NOTE: we could choose to duplicate the upgrade.Run unit tests for all of
	// this, but we choose to instead rely on end-to-end tests for most of this
	// functionality, and test only a few integration paths here.

	t.Run("when pg_upgrade --check fails it returns an error", func(t *testing.T) {
		execCommand = exectest.NewCommand(FailedMain)
		defer func() { execCommand = nil }()

		request := &idl.UpgradePrimariesRequest{
			SourceBinDir: "/old/bin",
			TargetBinDir: "/new/bin",
			DataDirPairs: pairs,
			CheckOnly:    true,
			UseLinkMode:  false,
		}
		err := UpgradePrimaries(tempDir, request, &spyMasterDataDirBackupTask{})
		if err == nil {
			t.Fatal("UpgradeSegments() returned no error")
		}

		// XXX it'd be nice if we didn't couple against a hardcoded string here,
		// but it's difficult to unwrap multierror with the new xerrors
		// interface.
		if !strings.Contains(err.Error(), "failed to check primary on host") ||
			!strings.Contains(err.Error(), "with content 1") {
			t.Errorf("error %q did not contain expected contents 'check primary on host' and 'content 1'",
				err.Error())
		}
	})

	t.Run("when pg_upgrade with no check fails it returns an error", func(t *testing.T) {
		execCommand = exectest.NewCommand(FailedMain)
		defer func() { execCommand = nil }()

		request := &idl.UpgradePrimariesRequest{
			SourceBinDir: "/old/bin",
			TargetBinDir: "/new/bin",
			DataDirPairs: pairs,
			CheckOnly:    false,
			UseLinkMode:  false}
		err := UpgradePrimaries(tempDir, request, &spyMasterDataDirBackupTask{})
		if err == nil {
			t.Fatal("UpgradeSegments() returned no error")
		}

		// XXX it'd be nice if we didn't couple against a hardcoded string here,
		// but it's difficult to unwrap multierror with the new xerrors
		// interface.
		if !strings.Contains(err.Error(), "failed to upgrade primary on host") ||
			!strings.Contains(err.Error(), "with content 1") {
			t.Errorf("error %q did not contain expected contents 'upgrade primary on host' and 'content 1'",
				err.Error())
		}
	})

	t.Run("it does not perform a copy of the master backup directory when using check mode", func(t *testing.T) {
		execCommand = exectest.NewCommand(Success)
		defer func() { execCommand = nil }()

		spyTask := &spyMasterDataDirBackupTask{}
		request := buildRequest(pairs)
		request.CheckOnly = true

		_ = UpgradePrimaries(tempDir, request, spyTask)

		result := spyTask.restoreCalls

		if len(result) != 0 {
			t.Errorf("recieved %v calls to rsync master data directory into segment data directory, want %v",
				len(result),
				0)
		}
	})

	t.Run("it grabs a copy of the master backup directory before running upgrade", func(t *testing.T) {
		execCommand = exectest.NewCommand(Success)
		defer func() { execCommand = nil }()

		spyTask := &spyMasterDataDirBackupTask{}

		request := buildRequest(pairs)
		request.MasterBackupDir = "/some/master/backup/dir"

		UpgradePrimaries(tempDir, request, spyTask)

		result := spyTask.restoreCalls

		if len(result) != len(pairs) {
			t.Errorf("recieved %v calls to rsync, want %v",
				len(result),
				len(pairs))
		}

		requestedFirstPair := pairs[0]
		firstRestoreCall := result[0]

		if firstRestoreCall.sourceDir != "/some/master/backup/dir" {
			t.Errorf("rsync source directory was %v, want %v",
				firstRestoreCall.sourceDir,
				"/some/master/backup/dir")
		}

		if firstRestoreCall.targetDir != requestedFirstPair.TargetDataDir {
			t.Errorf("rsync target directory was %v, want %v",
				firstRestoreCall.targetDir,
				requestedFirstPair.TargetDataDir)
		}
	})
}

func buildRequest(pairs []*idl.DataDirPair) *idl.UpgradePrimariesRequest {
	return &idl.UpgradePrimariesRequest{
		SourceBinDir:    "/old/bin",
		TargetBinDir:    "/new/bin",
		DataDirPairs:    pairs,
		CheckOnly:       false,
		UseLinkMode:     false,
		MasterBackupDir: "/some/master/backup/dir",
	}
}

type spyMasterDataDirBackupTask struct {
	restoreCalls []RestoreRequest
}

type RestoreRequest struct {
	sourceDir string
	targetDir string
}

func (c *spyMasterDataDirBackupTask) Restore(sourceDir, targetDir string) error {
	c.restoreCalls = append(c.restoreCalls, RestoreRequest{
		sourceDir: sourceDir,
		targetDir: targetDir,
	})

	return nil
}
