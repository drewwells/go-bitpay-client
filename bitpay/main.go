package main

import (
	"encoding/json"
	"fmt"

	bitpay "github.com/drewwells/go-bitpay-client"
	"github.com/kr/pretty"
)

func main() {
	resp := bitpay.Rates()
	// JSON is returned
	var v interface{}
	json.Unmarshal(resp, &v)
	fmt.Printf("% #v", pretty.Formatter(v))
}
