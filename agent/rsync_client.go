package agent

import "github.com/greenplum-db/gp-common-go-libs/cluster"

type rsyncClient struct{}

func (rsyncClient) Copy(sourceDir, targetDir string) error {
	return RestoreSegmentDataDir(targetDir, sourceDir, &cluster.GPDBExecutor{})
}

func NewRsyncClient() *rsyncClient {
	return &rsyncClient{}
}
