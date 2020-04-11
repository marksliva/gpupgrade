package agent

import (
	"context"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/upgrade"
)

// todo: add back the [...]
var postgresFiles = []string {"postgresql.conf", "PG_VERSION"}
var stateDirectoryFiles = []string {"config.json", "status.json"}

func (s *Server) DeleteStateDirectory(ctx context.Context, in *idl.DeleteStateDirectoryRequest) (*idl.DeleteStateDirectoryReply, error) {
	gplog.Info("got a request to delete the state directory from the hub")

	mErr := upgrade.DeleteDirectories([]string{in.Directory}, stateDirectoryFiles)
	return &idl.DeleteStateDirectoryReply{}, mErr.ErrorOrNil()
}

func (s *Server) DeleteDirectories(ctx context.Context, in *idl.DeleteDirectoriesRequest) (*idl.DeleteDirectoriesReply, error) {
	gplog.Info("got a request to delete data directories from the hub")

	mErr := upgrade.DeleteDirectories(in.Datadirs, postgresFiles)
	return &idl.DeleteDirectoriesReply{}, mErr.ErrorOrNil()
}
