package hub_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/greenplum-db/gpupgrade/greenplum"
	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"
)

func TestDeleteStateDirectories(t *testing.T) {
	c := hub.MustCreateCluster(t, []greenplum.SegConfig{
		{ContentID: -1, DbID: 0, Port: 25431, Hostname: "master", DataDir: "/data/qddir", Role: greenplum.PrimaryRole},
		{ContentID: -1, DbID: 1, Port: 25431, Hostname: "standby", DataDir: "/data/standby", Role: greenplum.MirrorRole},
		{ContentID: 0, DbID: 2, Port: 25432, Hostname: "sdw1", DataDir: "/data/dbfast1/seg1", Role: greenplum.PrimaryRole},
		{ContentID: 0, DbID: 6, Port: 35432, Hostname: "sdw1", DataDir: "/data/dbfast_mirror1/seg1", Role: greenplum.MirrorRole},
	})

	testhelper.SetupTestLogger() // initialize gplog

	t.Run("DeleteStateDirectories", func(t *testing.T) {
		t.Run("deletes state directories on all non-master hosts", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			expected := "/my/state_dir"

			sdw1Client := mock_idl.NewMockAgentClient(ctrl)
			sdw1Client.EXPECT().DeleteDirectory(
				gomock.Any(),
				&idl.DeleteDirectoryRequest{Directory: expected},
			).Return(&idl.DeleteDirectoryReply{}, nil)

			standbyClient := mock_idl.NewMockAgentClient(ctrl)
			standbyClient.EXPECT().DeleteDirectory(
				gomock.Any(),
				&idl.DeleteDirectoryRequest{Directory: expected},
			).Return(&idl.DeleteDirectoryReply{}, nil)

			masterClient := mock_idl.NewMockAgentClient(ctrl)
			// NOTE: we expect no call to the master

			agentConns := []*hub.Connection{
				{nil, sdw1Client, "sdw1", nil},
				{nil, standbyClient, "standby", nil},
				{nil, masterClient, "master", nil},
			}

			err := hub.DeleteStateDirectories(agentConns, c.MasterHostname())
			if err != nil {
				t.Errorf("unexpected err %#v", err)
			}
		})
	})

	//t.Run("DeleteDataDirectories", func(t *testing.T) {
	//	t.Run("returns error on failure", func(t *testing.T) {
	//		ctrl := gomock.NewController(t)
	//		defer ctrl.Finish()
	//
	//		sdw1Client := mock_idl.NewMockAgentClient(ctrl)
	//		sdw1Client.EXPECT().DeleteDirectories(
	//			gomock.Any(),
	//			gomock.Any(),
	//		).Return(&idl.DeleteDirectoriesReply{}, nil)
	//
	//		expected := errors.New("permission denied")
	//		sdw2ClientFailed := mock_idl.NewMockAgentClient(ctrl)
	//		sdw2ClientFailed.EXPECT().DeleteDirectories(
	//			gomock.Any(),
	//			gomock.Any(),
	//		).Return(nil, expected)
	//
	//		agentConns := []*hub.Connection{
	//			{nil, sdw1Client, "sdw1", nil},
	//			{nil, sdw2ClientFailed, "sdw2", nil},
	//		}
	//
	//		err := hub.DeleteDataDirectories(agentConns, c, false)
	//
	//		var multiErr *multierror.Error
	//		if !xerrors.As(err, &multiErr) {
	//			t.Fatalf("got error %#v, want type %T", err, multiErr)
	//		}
	//
	//		if len(multiErr.Errors) != 1 {
	//			t.Errorf("received %d errors, want %d", len(multiErr.Errors), 1)
	//		}
	//
	//		for _, err := range multiErr.Errors {
	//			if !xerrors.Is(err, expected) {
	//				t.Errorf("got error %#v, want %#v", expected, err)
	//			}
	//		}
	//	})
	//})
}
