package integrations_test

import (
	"errors"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"github.com/greenplum-db/gpupgrade/hub/upgradestatus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("prepare shutdown-clusters", func() {
	var (
		testExecutorOld *testhelper.TestExecutor
		testExecutorNew *testhelper.TestExecutor
	)

	BeforeEach(func() {
		testExecutorOld = &testhelper.TestExecutor{}
		testExecutorNew = &testhelper.TestExecutor{}
		source.Executor = testExecutorOld
		target.Executor = testExecutorNew
	})

	It("updates status PENDING, RUNNING then COMPLETE if successful", func() {
		Expect(cm.IsPending(upgradestatus.SHUTDOWN_CLUSTERS)).To(BeTrue())

		prepareShutdownClustersSession := runCommand("prepare", "shutdown-clusters")
		Eventually(prepareShutdownClustersSession).Should(Exit(0))

		Expect(testExecutorOld.NumExecutions).To(Equal(2))
		Expect(testExecutorOld.LocalCommands[0]).To(ContainSubstring("pgrep"))
		Expect(testExecutorOld.LocalCommands[1]).To(ContainSubstring(source.BinDir + "/gpstop -a"))

		Expect(testExecutorNew.NumExecutions).To(Equal(2))
		Expect(testExecutorNew.LocalCommands[0]).To(ContainSubstring("pgrep"))
		Expect(testExecutorNew.LocalCommands[1]).To(ContainSubstring(target.BinDir + "/gpstop -a"))

		Expect(cm.IsComplete(upgradestatus.SHUTDOWN_CLUSTERS)).To(BeTrue())
	})

	It("updates status to FAILED if it fails to run", func() {
		Expect(cm.IsPending(upgradestatus.SHUTDOWN_CLUSTERS)).To(BeTrue())

		testExecutorOld.ErrorOnExecNum = 2
		testExecutorNew.ErrorOnExecNum = 2
		testExecutorOld.LocalError = errors.New("stop failed")
		testExecutorNew.LocalError = errors.New("stop failed")

		prepareShutdownClustersSession := runCommand("prepare", "shutdown-clusters")
		Eventually(prepareShutdownClustersSession).Should(Exit(0))

		Expect(testExecutorOld.NumExecutions).To(Equal(2))
		Expect(testExecutorOld.LocalCommands[0]).To(ContainSubstring("pgrep"))
		Expect(testExecutorOld.LocalCommands[1]).To(ContainSubstring(source.BinDir + "/gpstop -a"))
		Expect(testExecutorOld.NumExecutions).To(Equal(2))
		Expect(testExecutorNew.LocalCommands[0]).To(ContainSubstring("pgrep"))
		Expect(testExecutorNew.LocalCommands[1]).To(ContainSubstring(target.BinDir + "/gpstop -a"))
		Expect(cm.IsFailed(upgradestatus.SHUTDOWN_CLUSTERS)).To(BeTrue())
	})
})
