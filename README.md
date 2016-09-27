# softlayer-go

SoftLayer API Client for the Go language

## Introduction

This library contains a complete implementation of the SoftLayer API for client application development in the Go programming language. Code for each API data type and service method is pre-generated, using the SoftLayer API metadata endpoint as input, thus ensuring 100% coverage of the API right out of the gate.

The library was designed to feel as natural as possible for programmers familiar with other popular SoftLayer SDKs, and attempts to minimize unnecessary boilerplate and type assertions where possible.

## Usage

### Basic example:

Three easy steps:

```go
// 1. Create a `Session`
sess := session.New(username, apikey)

// 2. Get a service
accountService := services.GetAccountService(sess)

// 3. Invoke a method:
account, err := accountService.getObject()
```

### Sessions

In addition to the example above, sessions can also be created using values
set in the environment, or from the local config file:

```go
sess := session.New()
```

In this usage, the username, API key, and endpoint are read from specific environment
variables, then the local configuration file (~/.softlayer).  First match ends
the search:

* Username: `env.SL_USERNAME` || `env.SOFTLAYER_USERNAME` || `config.username`
* API Key: `env.SL_API_KEY` || `env.SOFTLAYER_API_KEY` || `config.api_key`
* Endpoint: `env.SL_ENDPOINT_URL` || `env.SOFTLAYER_ENDPOINT_URL` || `config.endpoint_url`

*Note:* Endpoint defaults to `https://api.softlayer.com/rest/v3` if not configured through any of the above methods

### Instance methods

To call a method on a specific instance, set the instance ID before making the call:

```go
service := services.GetUserCustomerService(sess)

service.Id(6786566).GetObject()
```

### Passing Parameters

All non-slice method parameters are passed as pointers. This is to allow for optional values to be omitted (by passing `nil`)

```go
guestId := 123456
userCustomerService.RemoveVirtualGuestAccess(&guestId)
```

For convenience, a set of helper functions exists that allocate storage for a literal and return a pointer. The above can be refactored to:

```go
userCustomerService.RemoveVirtualGuestAccess(sl.Int(123456))
```

### Using datatypes

A complete library of SoftLayer API data type structs exists in the `datatypes` package. Like method parameters, all non-slice members are declared as pointers. This has the advantage of permitting updates without re-sending the complete data structure (since `nil` values are omitted from the resulting JSON). Use the same set of helper functions to assist in populating individual members.

```go
// Get the Virtual_Guest service
service := services.GetVirtualGuestService(sess)

// Create an empty Virtual_Guest struct as a template
vGuestTemplate := datatypes.Virtual_Guest{}

// Set Creation values - use helpers from the sl package to set pointer values.
// Unset (nil) values are not sent
vGuestTemplate.Hostname = sl.String("sample")
vGuestTemplate.Domain = sl.String("example.com")
vGuestTemplate.MaxMemory = sl.Int(4096)
vGuestTemplate.StartCpus = sl.Int(1)
vGuestTemplate.Datacenter = &datatypes.Location{Name: sl.String("wdc01")}
vGuestTemplate.OperatingSystemReferenceCode = sl.String("UBUNTU_LATEST")
vGuestTemplate.LocalDiskFlag = sl.Bool(true)

newGuest, err := service.CreateObject(&vGuestTemplate)

// optional error checking...

// Print the ID of the new guest.  Don't forget to dereference
fmt.Printf("New guest %d created", *newGuest.Id)
```

### Object Masks, Filters, Result Limits

Object masks, object filters, and pagination (limit and offset) can be set
by calling the `Mask()`, `Filter()`, `Limit()` and `Offset()` service methods
prior to invoking an API method.

For example, to set an object mask and filter that will be applied to
the results of the Account.GetObject() method:

```go
accountService := services.GetAccountService(sess)

accountService.
	Mask("id;hostname").
	Filter(`{"virtualGuests":{"domain":{"operation":"example.com"}}}`).
	GetObject()
```

The mask and filter are applied to the current request only, and reset after the
method returns. To preserve these options for future requests, save the return value:

```go
accountServiceWithMaskAndFilter = accountService.Mask("id;hostname").
	Filter(`{"virtualGuests":{"domain":{"operation":"example.com"}}}`)
```

Result limits are specified as separate `Limit` and `Offset` values:

```go
accountSerice.
	Offset(100).      // start at the 100th element in the list
	Limit(25).        // limit to 25 results
	GetVirtualGuests()
```

#### Filter Builder

There is also a **filter builder** you can use to create a _Filter_ instead of writing out the raw string:

```go
// requires importing the filter package
accountServiceWithMaskAndFilter = accountService.
    Mask("id;hostname").
    Filter(filter.Path("virtualGuests.domain").Eq("example.com").Build())
```

You can also create a filter incrementally:

```go
// Create initial filters
filters := filter.New(
    filter.Path("virtualGuests.hostname").StartsWith("KM078"),
    filter.Path("virtualGuests.id").NotEq(12345),
)

// ....
// Later, append another filter
filters = append(filters, filter.Path("virtualGuests.domain").Eq("example.com"))

accountServiceWithMaskAndFilter = accountService.
    Mask("id;hostname").
    Filter(filters.Build())
```

Or you can build all those filters in one step:

```go
// Create initial filters
filters := filter.Build(
    filter.Path("virtualGuests.hostname").StartsWith("KM078"),
    filter.Path("virtualGuests.id").NotEq(12345),
    filter.Path("virtualGuests.domain").Eq("example.com"),
)

accountServiceWithMaskAndFilter = accountService.
    Mask("id;hostname").Filter(filters)
```

See _filter/filters.go_ for the full range of operations supported.
The file at _examples/filters.go_ will show additional examples.
Also, [this is a good article](https://sldn.softlayer.com/article/object-filters) that describes SoftLayer filters at length.

### Handling Errors

For any error that occurs within one of the SoftLayer API services, a custom
error type is returned, with individual fields that can be parsed separately.

```go
_, err := service.Id(0).      // invalid object ID
	GetObject()

if err != nil {
	// Note: type assertion is only necessary for inspecting individual fields
	apiErr := err.(sl.Error)
	fmt.Printf("API Error:")
	fmt.Printf("HTTP Status Code: %d\n", apiErr.StatusCode)
	fmt.Printf("API Code: %s\n", apiErr.Exception)
	fmt.Printf("API Error: %s\n", apiErr.Message)
}
```

Note that `sl.Error` implements the standard `error` interface, so it can
be handled like any other error, if the above granularity is not needed:

```go
_, err := service.GetObject()
if err != nil {
	fmt.Println("Error during processing: ", err)
}
```

### Session Options

To set a different endpoint (e.g., the backend network endpoint):

```go
session.Endpoint = "https://api.service.softlayer.com/rest/v3"
```

To enable debug output:

```go
session.Debug = true
```

### Password-based authentication

Password-based authentication (via requesting a token from the API) is
only supported when talking to the API using the XML-RPC transport protocol.

To use the XML-RPC protocol, simply specify an XML-RPC endpoint url:

```go
func main() {
    // Create a session specifying an XML-RPC endpoint url.
    sess := &session.Session{
        Endpoint: "https://api.softlayer.com/xmlrpc/v3",
    }

    // Get a token from the api using your username and password
    userService := services.GetUserCustomerService(sess)
    token, err := userService.GetPortalLoginToken(username, password, nil, nil)
    if err != nil {
        log.Fatal(err)
    }

    // Add user id and token to the session.
    sess.UserId = *token.UserId
    sess.AuthToken = *token.Hash

    // You have a complete authenticated session now.
    // Call any api from this point on as normal...
    keys, err := userService.Id(sess.UserId).GetApiAuthenticationKeys()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("API Key:", *keys[0].AuthenticationKey)
}
```

## Development

### Setup

To get _softlayer-go_:

```
go get github.com/softlayer/softlayer-go
```

### Build

```
make build
```

### Test

```
make test
```

### Updating dependencies

```
make update_deps
```

## Copyright

This software is Copyright (c) 2016 IBM Corp. See the bundled LICENSE file for more information.
