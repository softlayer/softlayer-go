package sl_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/softlayer/softlayer-go/sl"
)

var _ = Describe("Options Tests", func() {
	var sl_limit int
	var sl_id int
	var sl_mask string
	var sl_filter string
	var sl_offset int
	var sl_totalItems int
	Context("Option Setter Testing", func() {
		BeforeEach(func() {
			sl_limit = 10
			sl_id = 99999
			sl_mask = "mask[id, hostname]"
			sl_filter = `{"test":{"operation":"ok"}}`
			sl_offset = 0
			sl_totalItems = 0
		})
		It("Test GetRemainingAPICalls", func() {
			options := sl.Options{
				TotalItems: 1000,
				Limit:      &sl_limit,
			}
			result := options.GetRemainingAPICalls()
			Expect(result).To(Equal(99))
		})
		It("Test ValidateLimit", func() {
			options := sl.Options{}
			limit := options.ValidateLimit()
			Expect(limit).To(Equal(sl.DefaultLimit))
			Expect(*options.Limit).To(Equal(sl.DefaultLimit))
			options.SetLimit(123)
			Expect(*options.Limit).To(Equal(123))
		})
		It("Setter Tests", func() {
			options := sl.Options{}
			options.SetTotalItems(44)
			Expect(options.TotalItems).To(Equal(44))
			options.SetOffset(33)
			Expect(*options.Offset).To(Equal(33))
			options.SetLimit(22)
			Expect(*options.Limit).To(Equal(22))
		})
		It("Basic Setting Tests", func() {
			options := sl.Options{
				Id:         &sl_id,
				Mask:       sl_mask,
				Filter:     sl_filter,
				Limit:      &sl_limit,
				Offset:     &sl_offset,
				TotalItems: sl_totalItems,
			}
			Expect(options.Limit).To(Equal(&sl_limit))
		})

	})
})
