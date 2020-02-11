package hub_test

import (
	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"
	"github.com/greenplum-db/gpupgrade/step"
	"testing"
)

var mockStepRunCalled = false

type MockStep struct {
	name    string
	sender  idl.MessageSender // sends substep status messages
	store   step.Store             // persistent substep status storage
	streams step.OutStreamsCloser  // writes substep stdout/err
	err     error
}

func (s *MockStep) Finish() error {
	return nil
}

func (s *MockStep) Err() error {
	return nil
}

func (s *MockStep) AlwaysRun(substep idl.Substep, f func(step.OutStreams) error) {

}

func (mockStep *MockStep) Run(substep idl.Substep, f func(step.OutStreams) error) {
	mockStepRunCalled = true
}

func MockBeginStep(stateDir string, name string, sender idl.MessageSender) (step.StepInterface, error) {
	return &MockStep{}, nil
}

func TestFinalize(t *testing.T) {
	t.Run("calls the reconfigure ports substep", func(t *testing.T) {
		mockServer := &hub.Server{StateDir: "/not/a/real/dir", Config: &hub.Config{Source:nil,Target:nil}}
		mockStream := &mock_idl.MockCliToHub_FinalizeServer{}

		err := hub.FinalizeStep(mockServer, mockStream, func(string, string, idl.MessageSender) (step.StepInterface, error) {
			return MockBeginStep(mockServer.StateDir, "finalize", mockStream)
		})

		if err != nil {
			t.Errorf("Expected success. Got %#v", err)
		}

		if mockStepRunCalled != true {
			t.Errorf("Expected Run() to be called, but it wasn't.")
		}
	})
}
