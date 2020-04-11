package upgrade_test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/greenplum-db/gpupgrade/upgrade"
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

func TestTempDataDir(t *testing.T) {
	var id upgrade.ID

	cases := []struct {
		datadir        string
		segPrefix      string
		expectedFormat string // %s will be replaced with id.String()
	}{
		{"/data/seg-1", "seg", "/data/seg.%s.-1"},
		{"/data/master/gpseg-1", "gpseg", "/data/master/gpseg.%s.-1"},
		{"/data/seg1", "seg", "/data/seg.%s.1"},
		{"/data/seg1/", "seg", "/data/seg.%s.1"},
		{"/data/standby", "seg", "/data/standby.%s"},
	}

	for _, c := range cases {
		actual := upgrade.TempDataDir(c.datadir, c.segPrefix, id)
		expected := fmt.Sprintf(c.expectedFormat, id)

		if actual != expected {
			t.Errorf("TempDataDir(%q, %q, id) = %q, want %q",
				c.datadir, c.segPrefix, actual, expected)
		}
	}
}

func ExampleTempDataDir() {
	var id upgrade.ID

	master := upgrade.TempDataDir("/data/master/seg-1", "seg", id)
	standby := upgrade.TempDataDir("/data/standby", "seg", id)
	segment := upgrade.TempDataDir("/data/primary/seg3", "seg", id)

	fmt.Println(master)
	fmt.Println(standby)
	fmt.Println(segment)
	// Output:
	// /data/master/seg.AAAAAAAAAAA.-1
	// /data/standby.AAAAAAAAAAA
	// /data/primary/seg.AAAAAAAAAAA.3
}

func TestDeleteDataDirectories(t *testing.T) {
	testhelper.SetupTestLogger()

	dataDirectories := []string{"/data/dbfast_mirror1/seg1", "/data/dbfast_mirror2/seg2"}

	t.Run("successfully deletes the data directories if all required files exist for that directory", func(t *testing.T) {
		filesThatMustExist := []string{"postgres-file-1", "postgres-file-2"}
		actualDataDirectories := []string{}
		utils.System.RemoveAll = func(name string) error {
			actualDataDirectories = append(actualDataDirectories, name)
			return nil
		}

		expectedFilesStatCalls := []string{"/data/dbfast_mirror1/seg1/postgres-file-1",
			"/data/dbfast_mirror1/seg1/postgres-file-2",
			"/data/dbfast_mirror2/seg2/postgres-file-1",
			"/data/dbfast_mirror2/seg2/postgres-file-2",
		}
		actualFilesStatCalls := []string{}
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			actualFilesStatCalls = append(actualFilesStatCalls, name)

			return nil, nil
		}

		err := upgrade.DeleteDirectories(dataDirectories, filesThatMustExist)

		if !reflect.DeepEqual(actualDataDirectories, dataDirectories) {
			t.Errorf("got %s, want %s", actualDataDirectories, dataDirectories)
		}

		if !reflect.DeepEqual(actualFilesStatCalls, expectedFilesStatCalls) {
			t.Errorf("got %s, want %s", actualFilesStatCalls, expectedFilesStatCalls)
		}

		if err != nil {
			t.Errorf("unexpected error got %+v", err)
		}
	})

	t.Run("fails to open configuration files under segment data directory", func(t *testing.T) {
		expected := errors.New("permission denied")
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			return nil, expected
		}

		err := upgrade.DeleteDirectories(dataDirectories, []string{"a", "b"})

		var multiErr *multierror.Error
		if !xerrors.As(err, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err, multiErr)
		}

		if len(multiErr.Errors) != 4 {
			t.Errorf("received %d errors, want %d", len(multiErr.Errors), 4)
		}

		for _, err := range multiErr.Errors {
			if !xerrors.Is(err, expected) {
				t.Errorf("got error %#v, want %#v", expected, err)
			}
		}
	})

	t.Run("fails to remove one segment data directory", func(t *testing.T) {
		expected := errors.New("permission denied")
		expectedDataDirectories := []string{"/data/dbfast_mirror1/seg1", "/data/dbfast_mirror2/seg2"}
		actualDataDirectories := []string{}
		utils.System.RemoveAll = func(name string) error {
			actualDataDirectories = append(actualDataDirectories, name)
			if name == "/data/dbfast_mirror1/seg1" {
				return expected
			}
			return nil
		}

		statCalls := 0
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			statCalls++
			return nil, nil
		}

		err := upgrade.DeleteDirectories(dataDirectories, []string{"foo", "bar"})

		var multiErr *multierror.Error
		if !xerrors.As(err, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err, multiErr)
		}

		if len(multiErr.Errors) != 1 {
			t.Errorf("got %d errors, want %d", len(multiErr.Errors), 1)
		}

		if statCalls != 4 {
			t.Errorf("got %d stat calls, want 4", statCalls)
		}

		if !reflect.DeepEqual(actualDataDirectories, expectedDataDirectories) {
			t.Errorf("got %s, want %s", actualDataDirectories, expectedDataDirectories)
		}

		for _, err := range multiErr.Errors {
			if !xerrors.Is(err, expected) {
				t.Errorf("got error %#v, want %#v", expected, err)
			}
		}
	})
}
