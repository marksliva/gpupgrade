package commanders_test

import (
	"errors"

	"github.com/greenplum-db/gpupgrade/cli/commanders"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"

	"github.com/golang/mock/gomock"
	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/greenplum-db/gpupgrade/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("check config", func() {

	var (
		client *mock_idl.MockCliToHubClient
		ctrl   *gomock.Controller
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		client = mock_idl.NewMockCliToHubClient(ctrl)
	})

	AfterEach(func() {
		utils.System = utils.InitializeSystemFunctions()
		defer ctrl.Finish()
	})

	Describe("Execute", func() {
		It("prints out that configuration has been obtained from the segments"+
			" and saved in persistent store", func() {
			//testLogger, testStdout, testStderr, testLogfile := testutils.SetupTestLogger()
			testStdout, _, _ := testhelper.SetupTestLogger()

			fakeCheckConfigReply := &idl.CheckConfigReply{}
			client.EXPECT().CheckConfig(
				gomock.Any(),
				&idl.CheckConfigRequest{},
			).Return(fakeCheckConfigReply, nil)

			request := commanders.NewConfigChecker(client)
			err := request.Execute()
			Expect(err).To(BeNil())
			Eventually(testStdout).Should(gbytes.Say("Check config request is processed."))
		})

		It("prints out an error when connection cannot be established to the hub", func() {
			_, testStderr, _ := testhelper.SetupTestLogger()
			client.EXPECT().CheckConfig(
				gomock.Any(),
				&idl.CheckConfigRequest{},
			).Return(nil, errors.New("Force failure connection"))

			request := commanders.NewConfigChecker(client)
			err := request.Execute()
			Expect(err).ToNot(BeNil())
			Eventually(testStderr).Should(gbytes.Say("ERROR - gRPC call to hub failed"))

		})
	})

})
