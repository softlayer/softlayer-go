package hardware_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/helpers/hardware"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
	"github.com/softlayer/softlayer-go/sl"
	"testing"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helper Hardware Tests")
}

var _ = Describe("Helper Hardware Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	var options *sl.Options
	BeforeEach(func() {
		limit := 10
		slsession = &sessionfakes.FakeSLSession{}
		options = &sl.Options{
			Mask:   "mask[id,hostname]",
			Filter: "",
			Limit:  &limit,
		}
	})

	Context("GetHardwareIter Tests", func() {

		It("API call made properly", func() {
			_, err := hardware.GetHardwareIter(slsession, options)
			Expect(err).ToNot(HaveOccurred())
			Expect(slsession.DoRequestCallCount()).To(Equal(1))
		})
	})
})
