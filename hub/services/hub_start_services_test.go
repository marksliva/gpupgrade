package services_test

import (
	"errors"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/net/context"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	as "github.com/greenplum-db/gpupgrade/agent/services"
	"github.com/greenplum-db/gpupgrade/hub/services"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/testutils/exectest"
)

func gpupgrade_agent() {
}

func gpupgrade_agent_Errors() {
	os.Stderr.WriteString("could not find state-directory")
	os.Exit(1)
}

func init() {
	exectest.RegisterMains(
		gpupgrade_agent,
		gpupgrade_agent_Errors,
	)
}

func TestRestartAgent(t *testing.T) {
	testhelper.SetupTestLogger()
	listener := bufconn.Listen(1024 * 1024)
	agentServer := grpc.NewServer()
	defer agentServer.Stop()

	idl.RegisterAgentServer(agentServer, &as.AgentServer{})
	go func() {
		if err := agentServer.Serve(listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	hostnames := []string{"host1", "host2"}
	port := 6416
	stateDir := "/not/existent/directory"
	ctx := context.Background()

	services.SetExecCommand(exectest.NewCommand(gpupgrade_agent))
	defer services.ResetExecCommand()

	t.Run("does not start running agents", func(t *testing.T) {
		dialer := func(ctx context.Context, address string) (net.Conn, error) {
			return listener.Dial()
		}

		restartedHosts, err := services.RestartAgents(ctx, dialer, hostnames, port, stateDir)
		if err != nil {
			t.Errorf("returned %#v", err)
		}
		if len(restartedHosts) != 0 {
			t.Errorf("restarted hosts %v", restartedHosts)
		}
	})

	t.Run("only restarts down agents", func(t *testing.T) {
		expectedHost := "host1"

		dialer := func(ctx context.Context, address string) (net.Conn, error) {
			if strings.HasPrefix(address, expectedHost) { //fail connection attempts to expectedHost
				return nil, errors.New("ahhhhh")
			}

			return listener.Dial()
		}

		restartedHosts, err := services.RestartAgents(ctx, dialer, hostnames, port, stateDir)
		if err != nil {
			t.Errorf("returned %#v", err)
		}

		if len(restartedHosts) != 1 {
			t.Errorf("expected one host to be restarted, got %d", len(restartedHosts))
		}

		if restartedHosts[0] != expectedHost {
			t.Errorf("expected restarted host %s got: %v", expectedHost, restartedHosts)
		}
	})

	t.Run("returns an error when gpupgrade_agent fails", func(t *testing.T) {
		services.SetExecCommand(exectest.NewCommand(gpupgrade_agent_Errors))

		// we fail all connections here so that RestartAgents will run the
		//  (error producing) gpupgrade_agent_Errors
		dialer := func(ctx context.Context, address string) (net.Conn, error) {
			return nil, errors.New("ahhhh")
		}

		restartedHosts, err := services.RestartAgents(ctx, dialer, hostnames, port, stateDir)
		if err == nil {
			t.Errorf("expected restart agents to fail")
		}

		if merr, ok := err.(*multierror.Error); ok {
			if merr.Len() != 2 {
				t.Errorf("expected 2 errors, got %d", merr.Len())
			}

			var exitErr *exec.ExitError
			for _, err := range merr.WrappedErrors() {
				if !xerrors.As(err, &exitErr) || exitErr.ExitCode() != 1 {
					t.Errorf("expected exit code: 1 but got: %#v", err)
				}
			}
		}

		if len(restartedHosts) != 0 {
			t.Errorf("restarted hosts %v", restartedHosts)
		}
	})

}
