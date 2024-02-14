package virtual_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// "github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
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
	var sl_service services.Account
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
		sl_service = services.GetAccountService(slsession)
	})

	Context("GetVirtualGuestsIter Tests", func() {
		It("API call made properly", func() {
			_, err := virtual.GetVirtualGuestsIter(sl_service)
			Expect(err).ToNot(HaveOccurred())
			Expect(slsession.DoRequestCallCount()).To(Equal(1))
		})
	})
}) 