package agent

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/greenplum-db/gpupgrade/idl"
)

var sedCommand = exec.Command

func (s *Server) UpdateRecoveryConfs(ctx context.Context, request *idl.UpdateRecoveryConfsRequest) (*idl.UpdateRecoveryConfsReply, error) {
	err := UpdateRecoveryConfPorts(request)

	return &idl.UpdateRecoveryConfsReply{}, err
}

func UpdateRecoveryConfPorts(recoverConfInfos *idl.UpdateRecoveryConfsRequest) error {
	for _, recoveryConfInfo := range recoverConfInfos.RecoveryConfInfos {
		sedCommandString := fmt.Sprintf("sed -i'.bak' 's/port=%d/port=%d/' %s",
			recoveryConfInfo.TemporaryPort,
			recoveryConfInfo.SourcePort,
			filepath.Join(recoveryConfInfo.DataDir, "recovery.conf",
		))

		err := sedCommand("bash", "-c", sedCommandString).Run()
		if err != nil {
			return err
		}
	}

	return nil
}
