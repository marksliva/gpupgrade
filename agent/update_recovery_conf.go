package agent

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

var sedCommand = exec.Command

type RecoverConfInfo struct {
	TemporaryPort int
	SourcePort    int
	DataDir       string
}

func UpdateRecoveryConfPorts(recoverConfInfos []RecoverConfInfo) error {
	for _, recoverConfInfo := range recoverConfInfos {
		sedCommandString := fmt.Sprintf("sed -i'.bak' 's/port=%d/port=%d/' %s",
			recoverConfInfo.TemporaryPort,
			recoverConfInfo.SourcePort,
			filepath.Join(recoverConfInfo.DataDir, "recovery.conf",
		))

		err := sedCommand("bash", "-c", sedCommandString).Run()
		if err != nil {
			return err
		}
	}

	return nil
}
