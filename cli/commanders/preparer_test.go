package commanders_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/greenplum-db/gpupgrade/cli/commanders"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"

	"github.com/golang/mock/gomock"
	"github.com/greenplum-db/gp-common-go-libs/testhelper"

	"github.com/greenplum-db/gpupgrade/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("preparer", func() {

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

	Describe("VerifyConnectivity", func() {
		It("returns nil when hub answers PingRequest", func() {
			testhelper.SetupTestLogger()

			client.EXPECT().Ping(
				gomock.Any(),
				&idl.PingRequest{},
			).Return(&idl.PingReply{}, nil)

			preparer := commanders.Preparer{}
			err := preparer.VerifyConnectivity(client)
			Expect(err).To(BeNil())
		})

		It("returns err when hub doesn't answer PingRequest", func() {
			testhelper.SetupTestLogger()
			commanders.NumberOfConnectionAttempt = 1

			client.EXPECT().Ping(
				gomock.Any(),
				&idl.PingRequest{},
			).Return(&idl.PingReply{}, errors.New("not answering ping")).Times(commanders.NumberOfConnectionAttempt + 1)

			preparer := commanders.Preparer{}
			err := preparer.VerifyConnectivity(client)
			Expect(err).ToNot(BeNil())
		})
		It("returns success if Ping eventually answers", func() {
			testhelper.SetupTestLogger()

			client.EXPECT().Ping(
				gomock.Any(),
				&idl.PingRequest{},
			).Return(&idl.PingReply{}, errors.New("not answering ping"))

			client.EXPECT().Ping(
				gomock.Any(),
				&idl.PingRequest{},
			).Return(&idl.PingReply{}, nil)

			preparer := commanders.Preparer{}
			err := preparer.VerifyConnectivity(client)
			Expect(err).To(BeNil())
		})
	})

	Describe("DoInit", func() {
		var (
			sourceBinDir string = "/old/does/not/exist"
			targetBinDir string = "/new/does/not/exist"
			dir          string
		)

		BeforeEach(func() {
			var err error
			dir, err = ioutil.TempDir("", "")
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			os.RemoveAll(dir)
		})

		It("populates cluster configuration files with the parameters it is passed", func() {
			stateDir := filepath.Join(dir, "foo")
			sourceFilename := filepath.Join(stateDir, utils.SOURCE_CONFIG_FILENAME)
			targetFilename := filepath.Join(stateDir, utils.TARGET_CONFIG_FILENAME)
			err := commanders.DoInit(stateDir, sourceBinDir, targetBinDir)
			Expect(err).To(BeNil())

			source := &utils.Cluster{ConfigPath: filepath.Join(stateDir, utils.SOURCE_CONFIG_FILENAME)}
			err = source.Load()
			Expect(err).To(BeNil())
			Expect(sourceFilename).To(BeAnExistingFile())
			Expect(source.BinDir).To(Equal(sourceBinDir))

			target := &utils.Cluster{ConfigPath: filepath.Join(stateDir, utils.TARGET_CONFIG_FILENAME)}
			err = target.Load()
			Expect(err).To(BeNil())
			Expect(targetFilename).To(BeAnExistingFile())
			Expect(target.BinDir).To(Equal(targetBinDir))
		})

		It("errs out when the state dir already exists", func() {
			err := commanders.DoInit(dir, "/old/does/not/exist", "/new/does/not/exist")
			Expect(err).ToNot(BeNil())
		})
	})
})
