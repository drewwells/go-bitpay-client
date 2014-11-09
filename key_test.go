package go_bitpay_client

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestKeygen(t *testing.T) {

	priv, pub, err := Keygen()
	if err != nil {
		log.Fatal(err)
	}
	_, _ = priv, pub
	fmt.Printf("Private: %x\n", priv)
	fmt.Printf("Public: %x\n", pub)
	// The output is random, so needs a test validating the
	// pub and priv key are connected.
}

func TestSin(t *testing.T) {
	s := "a72e6ed895e9ec2621bcccb93a11b73100887bdbb684cbfa6241b6076c54d0b4"
	s = "a72e6ed895e9ec2621bcccb93a11b73100887bdbb684cbfa6241b6076c54d0b4"
	src, _ := hex.DecodeString(s)
	pub := PublicFromPrivate([]byte(src))
	sin := Sin(pub)
	fmt.Printf("Public: %x\n", pub)
	fmt.Printf("Sin: %x\n", sin)
}
