package mock_agent

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/idl"

	"google.golang.org/grpc"
)

type MockAgentServer struct {
	addr       net.Addr
	grpcServer *grpc.Server
	numCalls   int
	mu         sync.Mutex

	UpgradeConvertPrimarySegmentsRequest *idl.UpgradePrimariesRequest
	CreateSegmentDataDirRequest          *idl.CreateSegmentDataDirRequest

	Err chan error
}

// NewMockAgentServer starts a locally running Agent server and returns a struct
// that facilitates unit testing. It also returns a hub.Dialer that will
// redirect any outgoing connections to this mock server, as well as the port
// that the server is listening on.
//
// XXX Why is the Dialer type that we need for this agent defined in the hub
// package?
func NewMockAgentServer() (*MockAgentServer, hub.Dialer, int) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	mockServer := &MockAgentServer{
		addr:       lis.Addr(),
		grpcServer: grpc.NewServer(),
		Err:        make(chan error, 10000),
	}

	idl.RegisterAgentServer(mockServer.grpcServer, mockServer)

	go func() {
		mockServer.grpcServer.Serve(lis)
	}()

	// Target this running server during dial.
	port := lis.Addr().(*net.TCPAddr).Port
	dialer := func(ctx context.Context, _ string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
		authority := fmt.Sprintf("127.0.0.1:%d", port)
		return grpc.DialContext(ctx, authority, opts...)
	}

	return mockServer, dialer, port
}

func (m *MockAgentServer) CheckDiskSpace(context.Context, *idl.CheckSegmentDiskSpaceRequest) (*idl.CheckDiskSpaceReply, error) {
	m.increaseCalls()

	return &idl.CheckDiskSpaceReply{}, nil
}

func (m *MockAgentServer) UpgradePrimaries(ctx context.Context, in *idl.UpgradePrimariesRequest) (*idl.UpgradePrimariesReply, error) {
	m.increaseCalls()

	m.mu.Lock()
	defer m.mu.Unlock()
	m.UpgradeConvertPrimarySegmentsRequest = in

	var err error
	if len(m.Err) != 0 {
		err = <-m.Err
	}

	return &idl.UpgradePrimariesReply{}, err
}

func (m *MockAgentServer) CreateSegmentDataDirectories(ctx context.Context, in *idl.CreateSegmentDataDirRequest) (*idl.CreateSegmentDataDirReply, error) {
	m.increaseCalls()

	m.mu.Lock()
	defer m.mu.Unlock()
	m.CreateSegmentDataDirRequest = in

	var err error
	if len(m.Err) != 0 {
		err = <-m.Err
	}

	return &idl.CreateSegmentDataDirReply{}, err
}

func (m *MockAgentServer) StopAgent(ctx context.Context, in *idl.StopAgentRequest) (*idl.StopAgentReply, error) {
	return &idl.StopAgentReply{}, nil
}

func (m *MockAgentServer) Stop() {
	m.grpcServer.Stop()
}

func (m *MockAgentServer) increaseCalls() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.numCalls++
}

func (m *MockAgentServer) NumberOfCalls() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.numCalls
}
