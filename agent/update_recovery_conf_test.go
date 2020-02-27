package agent_test

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/greenplum-db/gpupgrade/agent"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/testutils/exectest"
)

func sedMain() {}

func sedFailed() {
	os.Exit(1)
}

func init() {
	exectest.RegisterMains(
		sedFailed,
		sedMain,
	)
}

func TestUpdateRecoveryConfPorts(t *testing.T) {
	agent.SetSedCommand(nil)

	defer func() {
		agent.SetSedCommand(exec.Command)
	}()

	t.Run("it replaces the temporary port with the source port", func(t *testing.T) {
		called := false
		expectedArgStrings := [][]string{
			{
				"-c", "sed -i'.bak' 's/port=1234/port=8000/' /tmp/datadirs/mirror1_upgrade/gpseg0/recovery.conf",
			},
			{
				"-c", "sed -i'.bak' 's/port=1235/port=8001/' /tmp/datadirs/mirror2_upgrade/gpseg1/recovery.conf",
			},
		}
		numberOfCalls := 0

		sedCommandWithVerifier := exectest.NewCommandWithVerifier(sedMain, func(path string, args ...string) {
			called = true

			if path != "bash" {
				t.Errorf(`got: %q want "bash"`, path)
			}

			expectedArgs := expectedArgStrings[numberOfCalls]

			if !reflect.DeepEqual(args, expectedArgs) {
				t.Errorf("got args %#v want %#v", args, expectedArgs)
			}
			numberOfCalls++
		})

		agent.SetSedCommand(sedCommandWithVerifier)

		recoveryConfInfos := &idl.UpdateRecoveryConfsRequest{RecoveryConfInfos: []*idl.RecoveryConfInfo{
			{TemporaryPort: 1234, SourcePort: 8000, DataDir: "/tmp/datadirs/mirror1_upgrade/gpseg0"},
			{TemporaryPort: 1235, SourcePort: 8001, DataDir: "/tmp/datadirs/mirror2_upgrade/gpseg1"},
		}}

		err := agent.UpdateRecoveryConfPorts(recoveryConfInfos)

		if err != nil {
			t.Errorf("got error %+v want no error", err)
		}

		if !called {
			t.Errorf("Expected sedCommand to be called, but it was not")
		}

		if numberOfCalls != len(expectedArgStrings) {
			t.Errorf("got %d calls want %d calls", numberOfCalls, len(expectedArgStrings))
		}
	})

	t.Run("when there is an error running the sed command it returns it", func(t *testing.T) {
		agent.SetSedCommand(exectest.NewCommand(sedFailed))
		recoveryConfInfos := &idl.UpdateRecoveryConfsRequest{RecoveryConfInfos: []*idl.RecoveryConfInfo{
			{TemporaryPort: 1234, SourcePort: 8000, DataDir: "/tmp/datadirs/mirror1_upgrade/gpseg0"},
		}}

		err := agent.UpdateRecoveryConfPorts(recoveryConfInfos)
		// todo: verify the exit error from sedFailed matches the error thrown
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}
