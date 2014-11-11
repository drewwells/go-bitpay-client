package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
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
	var r []reflect.Value
	for _, v := range flag.Args() {
		if addr, ok := funcs[v]; ok {
			f := reflect.ValueOf(addr)
			r = f.Call([]reflect.Value{})
			if r != nil && len(r) > 0 {
				resp = (r[0]).Bytes()
			} else {
				log.Fatal(r)
			}
		}
	}
	// JSON is returned
	var v interface{}
	if len(resp) > 0 {
		err := json.Unmarshal(resp, &v)
		if err != nil {
			// Assume this is a keygen call
			fmt.Printf("private key: %x\n", (r[0]).Bytes())
			fmt.Printf("public key:  %x\n", (r[1]).Bytes())
			fmt.Printf("sin:         %s\n", (r[2]).Bytes())
			fmt.Println("Save these to your .env file")
		}
		fmt.Printf("% #v", pretty.Formatter(v))
	}
}
