# Smartbonus package for Go language

This package provides Go implementation of Smartbonus api. Supported all public api of smartbonus.

With 100% test coverage.

## Installation

Use the `go` command:

	$ go get github.com/srostyslav/smartbonus-api

## Requirements

Smartbonus package tested against Go 1.13.

## Example

```go
package main

import (
    smartbonus "github.com/srostyslav/smartbonus-api"
    "fmt"
)

func main() {
    // Creating sb: panic on error
    // Ask about params smartbonus team
    sb := smartbonus.NewSmartBonus("your store id", "https://your.smartbonus.com/api/v2/")

    // Get smartbonus info about client
    if client, err := sb.GetClient("0555555555"); err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("Client <phone: %s; name: %s; balance: %.2f>\n", client.Phone, client.Name, client.Balance)
    }

    // see tests for more
}

```
