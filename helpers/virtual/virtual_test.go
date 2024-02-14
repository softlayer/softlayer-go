package virtual_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// "github.com/softlayer/softlayer-go/datatypes"
	
	"github.com/softlayer/softlayer-go/sl"
	"github.com/softlayer/softlayer-go/helpers/virtual"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
	"testing"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helper Virtual Tests")
}

var _ = Describe("Helper Virtual Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	var options *sl.Options
	BeforeEach(func() {
		limit := 10
		slsession = &sessionfakes.FakeSLSession{}
		options = &sl.Options{
			Mask: "mask[id,hostname]",
			Filter: "",
			Limit: &limit,
		}
	})

	Context("GetVirtualGuestsIter Tests", func() {

		It("API call made properly", func() {
			_, err := virtual.GetVirtualGuestsIter(slsession, options)
			Expect(err).ToNot(HaveOccurred())
			Expect(slsession.DoRequestCallCount()).To(Equal(1))
		})
	})
}) 