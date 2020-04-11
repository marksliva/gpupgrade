package agent_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/greenplum-db/gpupgrade/agent"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils"
)

func TestDeleteDirectories(t *testing.T) {
	server := agent.NewServer(agent.Config{
		Port:     -1,
		StateDir: "",
	})

	testhelper.SetupTestLogger()

	t.Run("successfully deletes the data directories if all required files exist for that directory", func(t *testing.T) {
		expectedDataDirectories := []string{"/data/dbfast_mirror1/seg1", "/data/dbfast_mirror2/seg2"}
		actualDataDirectories := []string{}
		utils.System.RemoveAll = func(name string) error {
			actualDataDirectories = append(actualDataDirectories, name)
			return nil
		}

		expectedFilesStatCalls := []string{"/data/dbfast_mirror1/seg1/postgresql.conf",
			"/data/dbfast_mirror1/seg1/PG_VERSION",
			"/data/dbfast_mirror2/seg2/postgresql.conf",
			"/data/dbfast_mirror2/seg2/PG_VERSION",
		}
		actualFilesStatCalls := []string{}
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			actualFilesStatCalls = append(actualFilesStatCalls, name)

			return nil, nil
		}

		req := &idl.DeleteDirectoriesRequest{Datadirs: expectedDataDirectories}
		_, err := server.DeleteDirectories(context.Background(), req)

		if !reflect.DeepEqual(actualDataDirectories, expectedDataDirectories) {
			t.Errorf("got %s, want %s", actualDataDirectories, expectedDataDirectories)
		}

		if !reflect.DeepEqual(actualFilesStatCalls, expectedFilesStatCalls) {
			t.Errorf("got %s, want %s", actualFilesStatCalls, expectedFilesStatCalls)
		}

		if err != nil {
			t.Errorf("unexpected error got %+v", err)
		}
	})
}

func TestDeleteStateDirectory(t *testing.T) {
	server := agent.NewServer(agent.Config{
		Port:     -1,
		StateDir: "",
	})

	testhelper.SetupTestLogger()

	t.Run("successfully deletes the data directories if all required files exist for that directory", func(t *testing.T) {
		directory := "/gpupgrade/my/state/dir"
		var actualDirectory string
		utils.System.RemoveAll = func(name string) error {
			actualDirectory = name
			return nil
		}

		expectedFilesStatCalls := []string{
			"/gpupgrade/my/state/dir/config.json",
			"/gpupgrade/my/state/dir/status.json",
		}
		actualFilesStatCalls := []string{}
		utils.System.Stat = func(name string) (os.FileInfo, error) {
			actualFilesStatCalls = append(actualFilesStatCalls, name)

			return nil, nil
		}

		req := &idl.DeleteStateDirectoryRequest{Directory: directory}
		_, err := server.DeleteStateDirectory(context.Background(), req)

		if !reflect.DeepEqual(actualDirectory, directory) {
			t.Errorf("got %s, want %s", actualDirectory, directory)
		}

		if !reflect.DeepEqual(actualFilesStatCalls, expectedFilesStatCalls) {
			t.Errorf("got %s, want %s", actualFilesStatCalls, expectedFilesStatCalls)
		}

		if err != nil {
			t.Errorf("unexpected error got %+v", err)
		}
	})
}
