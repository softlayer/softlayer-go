package sl_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/softlayer/softlayer-go/sl"
)

var _ = Describe("Error Tests", func() {
	var sl_err sl.Error
	var base_err error
	Context("Error Tests", func() {
		BeforeEach(func() {
			base_err = errors.New("TTTTT")
			sl_err = sl.Error{
				StatusCode: 0,
				Exception:  "",
				Message:    "",
				Wrapped:    base_err,
			}
		})
		It("Basic Error Tests", func() {
			Expect(sl_err.Error()).To(Equal("TTTTT"))
			sl_err.Wrapped = nil
			Expect(sl_err.Error()).To(Equal(""))
			sl_err.Exception = "AAA"
			Expect(sl_err.Error()).To(Equal("AAA: "))
			sl_err.Message = "BBB"
			Expect(sl_err.Error()).To(Equal("AAA: BBB "))
			sl_err.StatusCode = 99
			Expect(sl_err.Error()).To(Equal("AAA: BBB (HTTP 99)"))
		})

	})
})
