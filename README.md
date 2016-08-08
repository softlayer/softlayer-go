# gopherlayer

SoftLayer API Client for the Go language

## Introduction

This library contains a complete implementation of the SoftLayer API for client application development in the Go programming language. Code for each API data type and service method is pre-generated, using the SoftLayer API metadata endpoint as input, thus ensuring 100% coverage of the API right out of the gate.

The library was designed to feel as natural as possible for programmers familiar with other popular SoftLayer SDKs, and attempts to minimize unnecessary boilerplate and type assertions where possible.

## Usage

### Basic example:

Three easy steps:

```go
// 1. Create a `Session`
session := service.NewSession(username, apikey)

// 2. Get a service
accountService := session.GetAccountService()

// 3. Invoke a method:
account, err := accountService.getObject()
```

### Instance methods

To call a method on a specific instance, set the instance ID before making the call:

```go
service := session.GetUserCustomerService()

service.Id(6786566)

service.GetObject()
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
service := session.GetVirtualGuestService()

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

To set an object mask or filter:

```go
accountService := session.GetAccountService()

accountService.Mask("id;hostname")

accountService.Filter(
	`{"virtualGuests":{"domain":{"operation":"example.com"}}}`)
```

Result limits are specified as separate `Limit` and `Offset` values:

```go
accountSerice.Offset(100) // start at the 100th element in the list

accountService.Limit(25)  // limit to 25 results
```

For convenience, the above methods, as well as the `Id()` method seen earlier, can be chained:

```go
service.Id(123456).Mask("id;hostname").Filter("...").Limit(25).Offset(100)
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

## Copyright

This software is Copyright (c) 2016 IBM Corp. See the bundled LICENSE file for more information.
