package agent

import (
	"os/exec"
)

type rsyncClient struct{}

func (c *rsyncClient) Copy(sourceDir, targetDir string) {
	error := exec.
		Command("rsync", "--archive", "--delete", sourceDir+"/", targetDir).
		Run()

	if error != nil {
		print("nothing")
	}
}

func NewRsyncClient() *rsyncClient {
	return &rsyncClient{}
}
