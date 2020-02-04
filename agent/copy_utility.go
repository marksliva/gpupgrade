package agent

import "os/exec"

var rsyncExecCommand = exec.Command

func CopyWithRsync(sourceDir, targetDir string, excludedFiles []string) error {
	arguments := append([]string{
		"--archive", "--delete",
		sourceDir + "/", targetDir,
	}, makeExclusionList(excludedFiles)...)

	command := rsyncExecCommand("rsync", arguments...)
	return command.Run()
}

func makeExclusionList(excludedFiles []string) []string {
	var exclusions []string
	for _, excludedFile := range excludedFiles {
		exclusions = append(exclusions, "--exclude", excludedFile)
	}
	return exclusions
}
