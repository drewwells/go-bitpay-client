package go_bitpay_client

import (
	"fmt"
	"log"
	"testing"
)

func TestKeygen(t *testing.T) {

	pub, priv, err := Keygen()
	if err != nil {
		log.Fatal(err)
	}
	_ = priv
	fmt.Println(len(pub), pub)
}
