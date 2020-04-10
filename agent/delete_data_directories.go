package agent

import (
	"context"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/upgrade"
)
// todo: needs to go in its own file
func (s *Server) DeleteDirectory(ctx context.Context, in *idl.DeleteDirectoryRequest) (*idl.DeleteDirectoryReply, error) {
	return nil, nil
}

func (s *Server) DeleteDirectories(ctx context.Context, in *idl.DeleteDirectoriesRequest) (*idl.DeleteDirectoriesReply, error) {
	gplog.Info("got a request to delete data directories from the hub")

	mErr := upgrade.DeleteDataDirectories(in.Datadirs)
	return &idl.DeleteDirectoriesReply{}, mErr.ErrorOrNil()
}
