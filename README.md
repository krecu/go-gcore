# go-gcore
[![Build Status](https://travis-ci.org/dstdfx/go-gcore.svg?branch=master)](https://travis-ci.org/dstdfx/go-gcore)  

Unofficial Go library for accessing the G-Core CDN API

It contains some parts of common and reseller G-Core CDN API.

Official docs are outdated, but anyway it's good to take a look at:
- https://docs.gcorelabs.com/cdn/

## Authentication ##

Before you start you need to have a G-Core account, you can sign in [here](https://gcorelabs.com).

To use a common G-Core API (basic account) construct a new Common client, then use the various services on the client to
access different parts of the G-Core API. For example, authenticate as a common user and get account info:

```go
// Your registration credentials
authOpts := gcore.AuthOptions{
    Username: "username",
    Password: "password",
}

// Create a new Common client instance
client := gcore.NewCommonClient()

// Get a token
// Empty context just for the sake of simplicity here
if err := client.Authenticate(context.Background(), authOpts); err != nil {
    panic(err)
}
// Get account details
account, _, err := client.Account.Details(context.Background())
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", account)
```

To use a reseller G-Core API  construct a new Reseller client, then use the various services on the client to
access different parts of the G-Core Reseller API. For example, authenticate as a reseller and get a list of activated clients:

```go
// Your reseller registration credentials
authOpts := gcore.AuthOptions{
    Username: "username",
    Password: "password",
}

// Create a new Reseller client instance
client := gcore.NewResellerClient()

// Get a token
// Empty context just for the sake of simplicity here
if err := client.Authenticate(context.Background(), authOpts); err != nil {
    panic(err)
}

listOpts := gcore.ListOpts{
    Activated: true,
}

// Get a list of activated clients assigned to this reseller account
clients, _, err := client.Clients.List(context.Background(), listOpts)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", clients)
```


## License ##
This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
