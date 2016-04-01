package e2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var ccPmBinPath string

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)

	SynchronizedBeforeSuite(func() []byte {
		binPath, err := gexec.Build("github.com/glestaris/cc/pm")
		Expect(err).NotTo(HaveOccurred())

		return []byte(binPath)
	}, func(data []byte) {
		ccPmBinPath = string(data)
	})

	SynchronizedAfterSuite(func() {}, func() {
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "End-to-end Suite")
}
