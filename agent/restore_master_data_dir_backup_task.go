package agent

import "github.com/greenplum-db/gp-common-go-libs/cluster"

type masterDataDirBackupTask struct{}

func (masterDataDirBackupTask) Restore(sourceDir, targetDir string) error {
	return RestoreSegmentDataDir(targetDir, sourceDir, &cluster.GPDBExecutor{})
}

func NewMasterDataDirBackupTask() *masterDataDirBackupTask {
	return &masterDataDirBackupTask{}
}
