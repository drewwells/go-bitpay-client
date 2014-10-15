package main

import (
	"fmt"

	bitpay "github.com/drewwells/go-bitpay-client"
)

func main() {
	resp := bitpay.Token()
	fmt.Println(string(resp))
}
