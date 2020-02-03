package agent

import "os/exec"

type copyUtility struct{}

func (c copyUtility) Copy(sourceDir, targetDir string, excludedFiles []string) error {
	arguments := append([]string{
		"--archive", "--delete",
		sourceDir + "/", targetDir,
	}, makeExclusionList(excludedFiles)...)

	command := exec.Command("rsync", arguments...)
	return command.Run()
}

func makeExclusionList(excludedFiles []string) []string {
	var exclusions []string
	for _, excludedFile := range excludedFiles {
		exclusions = append(exclusions, "--exclude", excludedFile)
	}
	return exclusions
}

func NewCopyUtility() *copyUtility {
	return &copyUtility{}
}
