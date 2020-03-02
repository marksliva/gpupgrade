package agent

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/idl"
)

var sedCommand = exec.Command

func (s *Server) UpdateRecoveryConfs(ctx context.Context, request *idl.UpdateRecoveryConfsRequest) (*idl.UpdateRecoveryConfsReply, error) {
	var mErr multierror.Error
	for _, recoveryConfInfo := range request.RecoveryConfInfos {
		script := fmt.Sprintf("sed -i'.bak' 's/port=%d/port=%d/' %s",
			recoveryConfInfo.GetTargetPrimaryPort(),
			recoveryConfInfo.GetSourcePrimaryPort(),
			filepath.Join(recoveryConfInfo.GetTargetMirrorDataDir(), "recovery.conf"))

		err := sedCommand("bash", "-c", script).Run()
		if err != nil {
			mErr = *multierror.Append(&mErr, err)
		}
	}

	return &idl.UpdateRecoveryConfsReply{}, mErr.ErrorOrNil()
}
