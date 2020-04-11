package hub_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

func TestDeleteStateDirectories(t *testing.T) {
	testhelper.SetupTestLogger() // initialize gplog

	t.Run("DeleteStateDirectories", func(t *testing.T) {
		t.Run("deletes state directories on all non-master hosts", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			expected := utils.GetStateDir()

			sdw1Client := mock_idl.NewMockAgentClient(ctrl)
			sdw1Client.EXPECT().DeleteStateDirectory(
				gomock.Any(),
				&idl.DeleteStateDirectoryRequest{Directory: expected},
			).Return(&idl.DeleteStateDirectoryReply{}, nil)

			standbyClient := mock_idl.NewMockAgentClient(ctrl)
			standbyClient.EXPECT().DeleteStateDirectory(
				gomock.Any(),
				&idl.DeleteStateDirectoryRequest{Directory: expected},
			).Return(&idl.DeleteStateDirectoryReply{}, nil)

			masterClient := mock_idl.NewMockAgentClient(ctrl)
			// NOTE: we expect no call to the master

			agentConns := []*hub.Connection{
				{nil, sdw1Client, "sdw1", nil},
				{nil, standbyClient, "standby", nil},
				{nil, masterClient, "master", nil},
			}

			err := hub.DeleteStateDirectories(agentConns, "master")
			if err != nil {
				t.Errorf("unexpected err %#v", err)
			}
		})
	})

	t.Run("DeleteDataDirectories", func(t *testing.T) {
		t.Run("returns error on failure", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sdw1Client := mock_idl.NewMockAgentClient(ctrl)
			sdw1Client.EXPECT().DeleteStateDirectory(
				gomock.Any(),
				gomock.Any(),
			).Return(&idl.DeleteStateDirectoryReply{}, nil)

			expected := errors.New("permission denied")
			sdw2ClientFailed := mock_idl.NewMockAgentClient(ctrl)
			sdw2ClientFailed.EXPECT().DeleteStateDirectory(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, expected)

			agentConns := []*hub.Connection{
				{nil, sdw1Client, "sdw1", nil},
				{nil, sdw2ClientFailed, "sdw2", nil},
			}

			err := hub.DeleteStateDirectories(agentConns, "master")

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
	})
}
