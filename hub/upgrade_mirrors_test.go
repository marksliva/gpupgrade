package hub

import (
	"bytes"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

type greenplumStub struct {
	run func(utilityName string, arguments ...string) error
}

func (g *greenplumStub) Run(utilityName string, arguments ...string) error {
	return g.run(utilityName, arguments...)
}

func TestUpgradeMirrors(t *testing.T) {
	t.Run("streams the gpaddmirrors config file format", func(t *testing.T) {
		initializeConfig := InitializeConfig{
			Mirrors: []utils.SegConfig{{
				DbID:      3,
				ContentID: 0,
				Port:      234,
				Hostname:  "localhost",
				DataDir:   "/data/mirrors_upgrade/seg0",
				Role:      "m",
			}, {
				DbID:      4,
				ContentID: 1,
				Port:      235,
				Hostname:  "localhost",
				DataDir:   "/data/mirrors_upgrade/seg1",
				Role:      "m",
			}},
		}
		var out bytes.Buffer

		writeGpAddmirrorsConfig(&initializeConfig, &out)

		lines := []string{
			"0|localhost|234|/data/mirrors_upgrade/seg0",
			"1|localhost|235|/data/mirrors_upgrade/seg1",
		}

		expected := strings.Join(lines, "\n") + "\n"

		if out.String() != expected {
			t.Errorf("got %q want %q", out.String(), expected)
		}
	})

	t.Run("returns errors from provided write stream", func(t *testing.T) {
		conf := &InitializeConfig{Mirrors: []utils.SegConfig{
			{DbID: 3, ContentID: 0, Port: 234, Hostname: "localhost", DataDir: "/data/mirrors/seg0", Role: "m"},
		}}

		writer := &failingWriter{errors.New("ahhh")}

		err := writeGpAddmirrorsConfig(conf, writer)
		if !xerrors.Is(err, writer.err) {
			t.Errorf("returned error %#v, want %#v", err, writer.err)
		}
	})

	t.Run("runAddMirrors runs gpaddmirrors with the created config file", func(t *testing.T) {
		expectedFilepath := "/add/mirrors/config_file"
		runCalled := false

		stub := &greenplumStub{
			func(utilityName string, arguments ...string) error {
				runCalled = true

				expected := "gpaddmirrors"
				if utilityName != expected {
					t.Errorf("ran utility %q, want %q", utilityName, expected)
				}

				var fs flag.FlagSet

				filepath := fs.String("i", "", "")
				quietMode := fs.Bool("a", false, "")

				err := fs.Parse(arguments)
				if err != nil {
					t.Fatalf("error parsing arguments: %+v", err)
				}

				if *filepath != expectedFilepath {
					t.Errorf("got filepath %q, want %q", *filepath, expectedFilepath)
				}

				if !*quietMode {
					t.Errorf("missing -a flag")
				}
				return nil
			},
		}

		err := runAddMirrors(stub, expectedFilepath)
		if err != nil {
			t.Errorf("returned error %+v", err)
		}

		if !runCalled {
			t.Errorf("GreenplumRunner.Run() was not called")
		}
	})

	t.Run("runAddMirrors bubbles up errors from the utility", func(t *testing.T) {
		stub := new(greenplumStub)

		expected := errors.New("ahhhh")
		stub.run = func(_ string, _ ...string) error {
			return expected
		}

		actual := runAddMirrors(stub, "")
		if !xerrors.Is(actual, expected) {
			t.Errorf("returned error %#v, want %#v", actual, expected)
		}
	})

	t.Run("UpgradeMirrors writes the add mirrors config to the and runs add mirrors", func(t *testing.T) {
		stateDir := "/the/state/dir"
		expectedFilepath := filepath.Join(stateDir, "add_mirrors_config")
		runCalled := false
		readPipe, writePipe, err := os.Pipe()
		if err != nil {
			t.Errorf("error creating pipes %#v", err)
		}

		utils.System.Create = func(name string) (*os.File, error) {
			if name != expectedFilepath {
				t.Errorf("got filepath %q want %q", name, expectedFilepath)
			}
			if err != nil {
				return nil, err
			}
			return writePipe, nil
		}

		initializeConfig := InitializeConfig{
			Mirrors: []utils.SegConfig{{
				DbID:      3,
				ContentID: 0,
				Port:      234,
				Hostname:  "localhost",
				DataDir:   "/data/mirrors_upgrade/seg0",
				Role:      "m",
			}, {
				DbID:      4,
				ContentID: 1,
				Port:      235,
				Hostname:  "localhost",
				DataDir:   "/data/mirrors_upgrade/seg1",
				Role:      "m",
			}},
		}

		stub := greenplumStub{run: func(utilityName string, arguments ...string) error {
			runCalled = true

			expectedUtility := "gpaddmirrors"
			if utilityName != expectedUtility {
				t.Errorf("ran utility %q, want %q", utilityName, expectedUtility)
			}

			var fs flag.FlagSet

			actualFilepath := fs.String("i", "", "")
			quietMode := fs.Bool("a", false, "")

			err := fs.Parse(arguments)
			if err != nil {
				t.Fatalf("error parsing arguments: %+v", err)
			}

			if *actualFilepath != expectedFilepath {
				t.Errorf("got filepath %q want %q", *actualFilepath, expectedFilepath)
			}

			if !*quietMode {
				t.Errorf("missing -a flag")
			}
			return nil
		}}

		err = UpgradeMirrors(stateDir, 6000, &initializeConfig, &stub)

		if err != nil {
			t.Errorf("got unexpected error from UpgradeMirrors %#v", err)
		}

		expectedLines := []string{
			"0|localhost|234|/data/mirrors_upgrade/seg0",
			"1|localhost|235|/data/mirrors_upgrade/seg1",
		}

		expectedFileContents := strings.Join(expectedLines, "\n") + "\n"
		fileContents, _ := ioutil.ReadAll(readPipe)

		if expectedFileContents != string(fileContents) {
			t.Errorf("got file contents %q want %q", fileContents, expectedFileContents)
		}

		if !runCalled {
			t.Errorf("GreenplumRunner.Run() was not called")
		}
	})

	t.Run("UpgradeMirrors returns the error when create file path fails", func(t *testing.T) {
		expectedError := errors.New("i'm an error")
		utils.System.Create = func(name string) (file *os.File, err error) {
			return nil, expectedError
		}

		err := UpgradeMirrors("", 6000, &InitializeConfig{}, &greenplumStub{})
		if !xerrors.Is(err, expectedError) {
			t.Errorf("returned error %#v want %#v", err, expectedError)
		}
	})

	t.Run("UpgradeMirrors returns the error when writing and closing the config file fails", func(t *testing.T) {
		utils.System.Create = func(name string) (file *os.File, err error) {
			// A nil file will result in failure.
			return nil, nil
		}

		// We need at least one config entry to cause something to be written.
		conf := &InitializeConfig{Mirrors: []utils.SegConfig{
			{DbID: 3, ContentID: 0, Port: 234, Hostname: "localhost", DataDir: "/data/mirrors/seg0", Role: "m"},
		}}

		stub := new(greenplumStub)
		stub.run = func(_ string, _ ...string) error {
			t.Errorf("gpaddmirrors should not have been called")
			return nil
		}

		err := UpgradeMirrors("/state/dir", 6000, conf, stub)

		var merr *multierror.Error
		if !xerrors.As(err, &merr) {
			t.Fatalf("returned error %#v, want error type %T", err, merr)
		}

		if len(merr.Errors) != 2 {
			t.Errorf("expected exactly two errors")
		}

		for _, err := range merr.Errors {
			if !xerrors.Is(err, os.ErrInvalid) {
				t.Errorf("returned error %#v want %#v", err, os.ErrInvalid)
			}
		}
	})

	t.Run("UpgradeMirrors returns the error when running the command fails", func(t *testing.T) {
		_, writePipe, _ := os.Pipe()

		utils.System.Create = func(name string) (file *os.File, err error) {
			return writePipe, nil
		}

		expectedErr := errors.New("the error happened")

		stub := &greenplumStub{run: func(utilityName string, arguments ...string) error {
			return expectedErr
		}}

		err := UpgradeMirrors("/state/dir", 6000, &InitializeConfig{}, stub)
		if !xerrors.Is(err, expectedErr) {
			t.Errorf("returned error %#v want %#v", err, expectedErr)
		}
	})
}
