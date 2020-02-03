package agent

type masterDataDirBackupTask struct {
	copyUtility   CopyUtility
	excludedFiles []string
}

func (t *masterDataDirBackupTask) Restore(sourceDir, targetDir string) error {
	// TODO: return errors
	t.copyUtility.Copy(sourceDir, targetDir, t.excludedFiles)

	return nil
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
