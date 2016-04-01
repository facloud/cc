package e2e_test

import (
	"os/exec"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("End-to-end tests", func() {
	var (
		packageName string
		downloadURL string
		timeout     int
		sess        *gexec.Session
	)

	JustBeforeEach(func() {
		var err error
		sess, err = gexec.Start(
			exec.Command(ccPmBinPath,
				"-package", packageName,
				"-url", downloadURL,
				"-timeout", strconv.Itoa(timeout),
			), GinkgoWriter, GinkgoWriter,
		)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should run", func() {
		Eventually(sess).Should(gexec.Exit(0))
	})
})
