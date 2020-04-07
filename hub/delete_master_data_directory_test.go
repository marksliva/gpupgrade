package hub_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

func TestDeleteMasterDataDirectory(t *testing.T) {
	t.Run("it deletes the directory if it looks like a DB directory", func(t *testing.T) {
		masterDataDir := "/data/qddir/demoDataDir.Whe_etPewAw.-1"
		var actualDataDir string
		utils.System.RemoveAll = func(name string) error {
			actualDataDir = name
			return nil
		}

		expectedFilesStatCalls := []string{
			masterDataDir,
			filepath.Join(masterDataDir, "postgresql.conf"),
			filepath.Join(masterDataDir, "PG_VERSION"),
		}
		actualFilesStatCalls :=  []string{}
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			actualFilesStatCalls = append(actualFilesStatCalls, name)

			return nil, nil
		}

		err := hub.DeleteMasterDataDirectory(masterDataDir)

		if actualDataDir != masterDataDir {
			t.Errorf("got %s, want %s", actualDataDir, masterDataDir)
		}

		if !reflect.DeepEqual(actualFilesStatCalls, expectedFilesStatCalls) {
			t.Errorf("got %s, want %s", actualFilesStatCalls, expectedFilesStatCalls)
		}

		if err != nil {
			t.Errorf("unexpected error got %+v", err)
		}
	})

	t.Run("it does not error when the directory does not exist", func(t *testing.T) {
		masterDataDir := "/data/qddir/demoDataDir.Whe_etPewAw.-1"

		var actualDataDir string
		expected := os.ErrNotExist
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			actualDataDir = name
			return nil, expected
		}

		err := hub.DeleteMasterDataDirectory(masterDataDir)

		if actualDataDir != masterDataDir {
			t.Errorf("got %q want %q", actualDataDir, masterDataDir)
		}

		if err != nil {
			t.Errorf("got unexpected error %+v", err)
		}
	})

	t.Run("it returns an error when the directory exists, but the postgres files are missing", func(t *testing.T) {
		expected := os.ErrNotExist
		masterDataDir := "/data/qddir/demoDataDir.Whe_etPewAw.-1"

		expectedFilesStatCalls := []string{
			masterDataDir,
			filepath.Join(masterDataDir, "postgresql.conf"),
			filepath.Join(masterDataDir, "PG_VERSION"),
		}
		actualFilesStatCalls :=  []string{}
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			// todo: should we check the name instead?
			actualFilesStatCalls = append(actualFilesStatCalls, name)
			if len(actualFilesStatCalls) == 1 {
				return nil, nil
			}

			return nil, expected
		}

		err := hub.DeleteMasterDataDirectory(masterDataDir)

		if !reflect.DeepEqual(actualFilesStatCalls, expectedFilesStatCalls) {
			t.Errorf("got %s, want %s", actualFilesStatCalls, expectedFilesStatCalls)
		}

		var multiErr *multierror.Error
		if !xerrors.As(err, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err, multiErr)
		}

		if len(multiErr.Errors) != 2 {
			t.Errorf("received %d errors, want %d", len(multiErr.Errors), len(expected.Error()))
		}

		for _, err := range multiErr.Errors {
			if !xerrors.Is(err, expected) {
				t.Errorf("got error %#v, want %#v", expected, err)
			}
		}
	})

	t.Run("it returns an error when a stat on the directory fails with something other than ErrNotExist", func(t *testing.T) {
		masterDataDir := "/data/qddir/demoDataDir.Whe_etPewAw.-1"

		var actualDataDir string
		expected := os.ErrClosed
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			actualDataDir = name
			return nil, expected
		}

		err := hub.DeleteMasterDataDirectory(masterDataDir)

		if actualDataDir != masterDataDir {
			t.Errorf("got %s, want %s", actualDataDir, masterDataDir)
		}

		var multiErr *multierror.Error
		if !xerrors.As(err, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err, multiErr)
		}

		if len(multiErr.Errors) != 1 {
			t.Errorf("received %d errors, want %d", len(multiErr.Errors), 1)
		}

		for _, err := range multiErr.Errors {
			if !xerrors.Is(err, expected) {
				t.Errorf("got error %#v, want %#v", expected, err)
			}
		}
	})

	t.Run("it returns an error if deleting the directory errors", func(t *testing.T) {
		masterDataDir := "/data/qddir/demoDataDir.Whe_etPewAw.-1"
		expected := os.ErrPermission
		utils.System.RemoveAll = func(name string) error {
			return expected
		}

		utils.System.Stat = func(name string) (os.FileInfo, error) {
			return nil, nil
		}

		err := hub.DeleteMasterDataDirectory(masterDataDir)

		var multiErr *multierror.Error
		if !xerrors.As(err, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err, multiErr)
		}

		if len(multiErr.Errors) != 1 {
			t.Errorf("received %d errors, want %d", len(multiErr.Errors), 1)
		}

		for _, err := range multiErr.Errors {
			if !xerrors.Is(err, expected) {
				t.Errorf("got error %#v, want %#v", expected, err)
			}
		}
	})
}
