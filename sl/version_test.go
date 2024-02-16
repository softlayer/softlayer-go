package sl_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/softlayer/softlayer-go/sl"
)

var _ = Describe("Version Tests", func() {
	Context("Get Version Tests", func() {
		It("Setting Version", func() {
			version := sl.VersionInfo{
				Major: 1, Minor: 2, Patch: 3, Pre: "",
			}
			Expect(version.String()).To(Equal("v1.2.3"))
		})
		It("Pre Version", func() {
			version := sl.VersionInfo{
				Major: 1, Minor: 2, Patch: 3, Pre: "a",
			}
			Expect(version.String()).To(Equal("v1.2.3-a"))
		})
		It("Base Version", func() {
			version := sl.Version
			Expect(version.String()).Should(HavePrefix("v"))
		})
	})
})
