package services_test

import (
	"errors"
	"time"

	"github.com/greenplum-db/gpupgrade/idl"

	"github.com/golang/mock/gomock"

	"github.com/greenplum-db/gpupgrade/hub/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("hub pings agents test", func() {
	var (
		pingerManager *services.PingerManager
	)

	BeforeEach(func() {
		pingerManager = &services.PingerManager{
			RPCClients:       []services.ClientAndHostname{{Client: client, Hostname: "doesnotexist"}},
			NumRetries:       10,
			PauseBeforeRetry: 1 * time.Millisecond,
		}
	})

	Describe("PingAllAgents", func() {
		It("grpc calls succeed, all agents are running", func() {
			client.EXPECT().PingAgents(
				gomock.Any(),
				&idl.PingAgentsRequest{},
			).Return(&idl.PingAgentsReply{}, nil)

			err := pingerManager.PingAllAgents()
			Expect(err).To(BeNil())
		})

		It("grpc calls fail, not all agents are running", func() {
			client.EXPECT().PingAgents(
				gomock.Any(),
				&idl.PingAgentsRequest{},
			).Return(&idl.PingAgentsReply{}, errors.New("call to agent fail"))

			err := pingerManager.PingAllAgents()
			Expect(err).To(MatchError("call to agent fail"))
		})

		It("grpc calls succeed, only one ping", func() {
			client.EXPECT().PingAgents(
				gomock.Any(),
				&idl.PingAgentsRequest{},
			).Return(&idl.PingAgentsReply{}, nil)

			err := pingerManager.PingPollAgents()
			Expect(err).ToNot(HaveOccurred())
		})

		It("grpc calls fail, ping timeout exceeded", func() {
			for i := 0; i < pingerManager.NumRetries; i++ {
				client.EXPECT().PingAgents(
					gomock.Any(),
					&idl.PingAgentsRequest{},
				).Return(&idl.PingAgentsReply{}, errors.New("call to agent fail"))
			}

			err := pingerManager.PingPollAgents()
			Expect(err).To(MatchError("call to agent fail"))
		})
	})
})
