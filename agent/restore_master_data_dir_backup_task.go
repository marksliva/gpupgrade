package agent

type masterDataDirBackupTask struct {
	copyUtility   CopyUtility
	excludedFiles []string
}

func (t *masterDataDirBackupTask) Restore(sourceDir, targetDir string) error {
	return t.copyUtility.Copy(sourceDir, targetDir, t.excludedFiles)
}

func NewMasterDataDirBackupTask(
	copyUtility CopyUtility,
	excludedFiles []string,
) *masterDataDirBackupTask {
	return &masterDataDirBackupTask{
		copyUtility,
		excludedFiles,
	}
}
