# softlayer-go

[![Build Status](https://github.com/softlayer/softlayer-go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/softlayer/softlayer-go/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/softlayer/softlayer-go?status.svg)](https://godoc.org/github.com/softlayer/softlayer-go)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Coverage Status](https://coveralls.io/repos/github/softlayer/softlayer-go/badge.svg?branch=master)](https://coveralls.io/github/softlayer/softlayer-go?branch=master)

The Official and Complete SoftLayer API Client for Golang (the Go programming language).

## Introduction

This library contains a complete implementation of the SoftLayer API for client application development in the Go programming language. Code for each API data type and service method is pre-generated, using the SoftLayer API metadata endpoint as input, thus ensuring 100% coverage of the API right out of the gate.

It was designed to feel as natural as possible for programmers familiar with other popular SoftLayer SDKs, and attempts to minimize unnecessary boilerplate and type assertions where possible.

## Usage

### Basic example:

Three easy steps:

```go
// 1. Create a session
sess := session.New(username, apikey)

// 2. Get a service
accountService := services.GetAccountService(sess)

// 3. Invoke a method:
account, err := accountService.GetObject()
```

[More examples](https://github.com/softlayer/softlayer-go/tree/master/examples)

### Sessions

In addition to the example above, sessions can also be created using values
set in the environment, or from the local configuration file (i.e. ~/.softlayer):

```go
sess := session.New()
```

In this usage, the username, API key, and endpoint are read from specific environment
variables, then the local configuration file (i.e. ~/.softlayer).  First match ends
the search:

* _Username_
	1. environment variable `SL_USERNAME`
	1. environment variable `SOFTLAYER_USERNAME`
	1. local config `username`.
* _API Key_
	1. environment variable `SL_API_KEY`
	1. environment variable `SOFTLAYER_API_KEY`
	1. local config `api_key`.
* _Endpoint_
	1. environment variable `SL_ENDPOINT_URL`
	1. environment variable `SOFTLAYER_ENDPOINT_URL`
	1. local config `endpoint_url`.
* _Timeout_
	1. environment variable `SL_TIMEOUT`
	1. environment variable `SOFTLAYER_TIMEOUT`
	1. local config `timeout`.

*Note:* Endpoint defaults to `https://api.softlayer.com/rest/v3` if not configured through any of the above methods. Timeout defaults to 120 seconds.

Example of the **~/.softlayer** local configuration file:
```
[softlayer]
username = <your username>
api_key = <your api key>
endpoint_url = <optional>
timeout = <optional>
```

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
package main

import (
	"fmt"
	"log"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

func main() {
	sess := session.New() // See above for details about creating a new session

	// Get the Virtual_Guest service
	service := services.GetVirtualGuestService(sess)

	// Create a Virtual_Guest struct as a template
	vGuestTemplate := datatypes.Virtual_Guest{
		// Set Creation values - use helpers from the sl package to set pointer values.
		// Unset (nil) values are not sent
		Hostname:                     sl.String("sample"),
		Domain:                       sl.String("example.com"),
		MaxMemory:                    sl.Int(4096),
		StartCpus:                    sl.Int(1),
		Datacenter:                   &datatypes.Location{Name: sl.String("wdc01")},
		OperatingSystemReferenceCode: sl.String("UBUNTU_LATEST"),
		LocalDiskFlag:                sl.Bool(true),
	}

	// Tell the API to create the virtual guest
	newGuest, err := service.CreateObject(&vGuestTemplate)
	// optional error checking...
	if err != nil {
		log.Fatal(err)
	}

	// Print the ID of the new guest.  Don't forget to dereference
	fmt.Printf("New guest %d created", *newGuest.Id)
}
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
accountService.
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

By default, the debug output is sent to standard output. You can customize this by setting up your own logger:

```go
import "github.com/softlayer/softlayer-go/session"

session.Logger = log.New(os.Stderr, "[CUSTOMIZED] ", log.LstdFlags)
```

You can also tell the session to retry the api requests if there is a timeout error:

```go
// Specify how many times to retry the request, the request timeout, and how much time to wait
// between retries.
services.GetVirtualGuestService(
	sess.SetTimeout(900).SetRetries(2).SetRetryWait(3)
).GetObject(...)
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

### IAM authentication

This library supports [IBM's IAM Authentication](https://cloud.ibm.com/docs/account?topic=account-iamoverview) (used by the `ibmcloud` cli for example). You will want to set the `IAMToken` and `IAMRefreshToken` properties on the session to make use of it.


```go
token := "Bearer aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa..."
refreshToken := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb..."
sess := &session.Session{
	Endpoint:         "https://api.softlayer.com/rest/v3.1",
	Debug:            true,
	Timeout:          90,
	IAMToken:         token,
	IAMRefreshToken:  refreshToken
}
```

You can bypass automatic IAM refresh by either not setting the `IAMRefreshToken` property, or by manually configuring the `TransportHandler`

```go
token := "Bearer aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa..."
handler := &session.RestTransport{}
sess := &session.Session{
	Endpoint:         "https://api.softlayer.com/rest/v3.1",
	Debug:            true,
	Timeout:          90,
	IAMToken:         token,
	TransportHandler  handler,
}
```

If you want to be able to record the new tokens in a config file (or elsewhere), you can configure an `IAMUpdaters` Observer which will take in as arguments the new tokens, allowing you to save them somewhere for reuse.

```go
type MyIamUpdater struct {
	debug bool
}
// This function is where you can configure the logic to save these new tokens somewhere
func (iamupdater *MyIamUpdater) Update(token string, refresh string) {
	fmt.Printf("[DEBUG] New Token: %s\n", token)
	fmt.Printf("[DEBUG] New Refresh Token: %s\n", refresh)
}
updater := &MyIamUpdater{debug: false}
token := "Bearer aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa..."
refreshToken := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb..."
sess := &session.Session{
	Endpoint:         "https://api.softlayer.com/rest/v3.1",
	IAMToken:         token,
	IAMRefreshToken  refreshToken,
}
sess.AddIAMUpdater(updater)
```
You can add multiple Updaters as well, they will be called in the order they are added. `MyIamUpdater.Update(token, refresh)` in this example will only be called when the token is actually refreshed.

## Running Examples

The [Examples](https://github.com/softlayer/softlayer-go/tree/master/examples) directory has a few rough examples scripts that can help you get started developing with this library.

The Examples have their own modules since there are a few dependencies here not needed in the normal softlayer-go repo.
The following will get you started:

```bash
cd examples
go mod vendor
go run main.go  --help
```

The examples use [SPF13's Cobra Command](https://github.com/spf13/cobra) library to build little command examples, which should be easy enough to 
expand upon if needed. The [SoftLayer-CLI](https://github.com/softlayer/softlayer-cli) project is also a good example of how to use the softlayer-go library.


## Development

### Setup

To get _softlayer-go_:

```
go get github.com/softlayer/softlayer-go/...
```

### Build

```
make
```

```bash
gofmt -w `find . -name '*.go' | grep -v vendor`
go vet -all $(go list ./... | grep -v datatypes)
go get -d -v ./...
go build ./...

```

### Test

```
make test
```

```bash
gofmt -w `find . -name '*.go' | grep -v vendor`
go generate ./session
go vet -all $(go list ./... | grep -v datatypes)
go mod vendor
go test $(go list ./... | grep -v '/vendor/') -timeout=30s -coverprofile coverage.out
```

### Updating dependencies

```
make update_deps
```

### Generate
This downloads the API definitions and creates go files for them.

(with make)
```bash
make generate
```

(manually)
```bash
go run ./tools generate
```


## Copyright

This software is Copyright (c) 2016-2024 IBM Corp. See the bundled LICENSE file for more information.
