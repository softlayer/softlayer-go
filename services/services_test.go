package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/datatypes"
	"testing"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Tests")
}

// This is required for a few of the Place order API methods that take in an interface
func GetOrderContainer() *datatypes.Container_Product_Order {
	ComplexType := "Test_Complex_Type"
	OrderContainer := &datatypes.Container_Product_Order{ComplexType: &ComplexType}
	return OrderContainer
}
