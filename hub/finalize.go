package hub

import (
	"fmt"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/greenplum-db/gpupgrade/step"
	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/idl"
)

func (s *Server) Finalize(_ *idl.FinalizeRequest, stream idl.CliToHub_FinalizeServer) (err error) {
	return FinalizeStep(s, stream, func(string, string, idl.MessageSender) (step.StepInterface, error) {
		return BeginStep(s.StateDir, "finalize", stream)
	})
}

func FinalizeStep(s *Server, stream idl.CliToHub_FinalizeServer, beginStepFunc func(string, string, idl.MessageSender) (step.StepInterface, error)) error {
	st, err := beginStepFunc(s.StateDir, "finalize", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := st.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("finalize: %s", err))
		}
	}()

	st.Run(idl.Substep_RECONFIGURE_PORTS, func(stream step.OutStreams) error {
		return s.ReconfigurePorts(stream)
	})

	return st.Err()
}
