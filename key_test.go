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
	src, _ := hex.DecodeString(s)
	pub := PublicFromPrivate([]byte(src))
	//pubstr := "02F840A04114081690223B7069071A70D6DABB891763B638CC20C7EC3BD58E6C86"

	compressed := append([]byte{0x04}, pub[1:]...)[:33]
	fmt.Println("Compressed: 025C6E29280398E8706A65A094391DB956C219B2B9E3950D8967BC5B029285B5BC")
	fmt.Printf("Compressed: %X\n", compressed)
	//pubstr = "025C6E29280398E8706A65A094391DB956C219B2B9E3950D8967BC5B029285B5BC"
	dec, _ := hex.DecodeString(string(compressed))
	sin := Sin(dec)
	_ = sin
	fmt.Printf("Public: %x\n", pub)
	fmt.Printf("Sin:    %x\n", sin)
}
