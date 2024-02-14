package generator_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/softlayer/softlayer-go/generator"
	"os"
)

var metaMinPath = "../sldnSample.json"
var _ = Describe("Minimal Tests", func() {
	var outDir string

	BeforeEach(func() {
		var err error
		outDir, err = os.MkdirTemp("", "softlayer-go-test*")
		Expect(err).NotTo(HaveOccurred())
		err = os.Mkdir(fmt.Sprintf("%s/services", outDir), 0750)
		Expect(err).NotTo(HaveOccurred())
		err = os.Mkdir(fmt.Sprintf("%s/datatypes", outDir), 0750)
		Expect(err).NotTo(HaveOccurred())
	})
	Context("Testing SLDN Generation", func() {
		Context("GenerateAPI From File", func() {
			It("Generates SLDN from test file", func() {
				GenerateAPI(metaMinPath, outDir)
				Expect(fmt.Sprintf("%s/services/account.go", outDir)).Should(BeAnExistingFile())
				Expect(fmt.Sprintf("%s/services/account_test.go", outDir)).Should(BeAnExistingFile())
				Expect(fmt.Sprintf("%s/datatypes/virtual.go", outDir)).Should(BeAnExistingFile())
			})
		})
	})
	AfterEach(func() {
		Expect(os.RemoveAll(outDir)).To(Succeed())
	})
})
