package generator_test

import (
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"reflect"
	"testing"
)

// Tests for each service/method that follows special case logic during code
// generation

func TestVirtualDiskImageType(t *testing.T) {
	expected := "*datatypes.Virtual_Disk_Image_Type"
	actual := reflect.TypeOf(datatypes.Virtual_Guest_Block_Device_Template_Group{}.ImageType).String()

	if actual != expected {
		t.Errorf("Expect type of Virtual_Guest_Block_Device_Template_Group.ImageType to be %s, but was %s", expected, actual)
	}
}

func TestDomainResourceRecordPropertiesCopiedToBaseType(t *testing.T) {
	fields := []string{"IsGatewayAddress", "Port", "Priority", "Protocol", "Service", "Weight"}
	for _, field := range fields {
		if _, ok := reflect.TypeOf(datatypes.Dns_Domain_ResourceRecord{}).FieldByName(field); !ok {
			t.Errorf("Expect property %s not found for datatypes.Dns_Domain_ResourceRecord", field)
		}
	}
}

func TestVoidPatchedReturnTypes(t *testing.T) {
	tests := map[interface{}]string{
		services.Network_Application_Delivery_Controller_LoadBalancer_Service{}:       "DeleteObject",
		services.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer{}: "DeleteObject",
		services.Network_Application_Delivery_Controller{}:                            "DeleteLiveLoadBalancerService",
	}

	for service, method := range tests {
		reflectedService := reflect.TypeOf(service)
		reflectedMethod, _ := reflectedService.MethodByName(method)
		if reflectedMethod.Type.NumOut() > 1 {
			t.Errorf("Expect %s() to have only one (error) return value, but multiple values found", reflectedService.String()+method)
		}
	}
}

func TestPlaceOrder(t *testing.T) {
	services := []interface{}{
		services.Billing_Order_Quote{},
		services.Product_Order{},
	}

	methods := []string{"PlaceOrder", "VerifyOrder", "PlaceQuote"}
	for _, service := range services {
		serviceType := reflect.TypeOf(service)
		for _, methodName := range methods {
			method, _ := serviceType.MethodByName(methodName)
			argType := method.Type.In(1).String()
			if argType != "interface {}" {
				t.Errorf(
					"Expect %s.%s() to accept interface {} as parameter, but %s found instead",
					serviceType.String(),
					methodName,
					argType,
				)
			}
		}
	}
}
