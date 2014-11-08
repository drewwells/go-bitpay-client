package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"

	bitpay "github.com/drewwells/go-bitpay-client"
	"github.com/kr/pretty"
)

func main() {
	flag.Parse()
	var resp []byte

	// Create a list of available commands and if found call them
	funcs := map[string]interface{}{
		"token":   bitpay.Token,
		"invoice": bitpay.Invoice,
		"bills":   bitpay.Bills,
		"bill":    bitpay.Bill,
		"rates":   bitpay.Rates,
		"keygen":  bitpay.Keygen,
	}

	for _, v := range flag.Args() {
		if addr, ok := funcs[v]; ok {
			f := reflect.ValueOf(addr)
			r := f.Call([]reflect.Value{})
			if r != nil && len(r) > 0 {
				resp = (r[0]).Bytes()
			}
		}
	}
	// JSON is returned
	var v interface{}
	if len(resp) > 0 {
		json.Unmarshal(resp, &v)
		fmt.Printf("% #v", pretty.Formatter(v))
	}
}
