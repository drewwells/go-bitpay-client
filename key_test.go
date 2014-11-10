package go_bitpay_client

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

var TESTPRIV = []byte("1f0dd50dd70a6fce3fbe81b922d89610742b37aa1eae9f82ffecb5571dcebfa4")

func TestKeygen(t *testing.T) {

	priv, pub, err := Keygen()
	if err != nil {
		log.Fatal(err)
	}
	_, _ = priv, pub

	// fmt.Printf("Private: %x\n", priv)
	// fmt.Printf("Public: %x\n", pub)

	// The output is random, so needs a test validating the
	// pub and priv key are connected.
}

func BenchmarkKeygen(b *testing.B) {
	priv := make([]byte, 32)
	pub := make([]byte, 33)
	for n := 0; n < b.N; n++ {
		priv, pub, _ = Keygen()
	}
	dummy, dummy = priv, pub
}

func TestPublicFromPrivate(t *testing.T) {
	full := hex.EncodeToString(PublicFromPrivate(TESTPRIV, false))
	comp := hex.EncodeToString(PublicFromPrivate(TESTPRIV, true))

	eful := "047dc432eef530a7d46066921af445ac67ae3147e93131910c1578c813af06375ef9b0e9987c243f0a78df60eac6f9c777f2ba66e57936a9bb1dccb76fead0c4a0"
	ecom := "027dc432eef530a7d46066921af445ac67ae3147e93131910c1578c813af06375e"

	if eful != full {
		t.Errorf("got:\n%s\nwant:\n%s\n",
			full, eful)
	}
	if ecom != comp {
		t.Errorf("got:\n%s\nwant:\n%s\n",
			comp, ecom)
	}
}

var dummy []byte

func BenchmarkPublicFromPrivate_Compressed(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		bs = PublicFromPrivate(TESTPRIV, true)
	}
	dummy = bs
}

func BenchmarkPublicFromPrivate_Uncompressed(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		bs = PublicFromPrivate(TESTPRIV, false)
	}
	dummy = bs
}
func ExampleSin() {
	txt := make([]byte, hex.DecodedLen(len(TESTPRIV)))
	hex.Decode(txt, TESTPRIV)
	pub := PublicFromPrivate(txt, false)
	pubcomp := PublicFromPrivate(txt, true)

	sin := Sin(pubcomp)

	fmt.Printf("Public:  %x\n", string(pub))
	fmt.Printf("PubComp: %x\n", string(pubcomp))
	fmt.Printf("Sin:     %s\n", string(sin))

	// Output:
	// Public:  04c6072aad509c88edad53756abc01f00f9e3feeb08fd748b4081964bae97e253214b138fb752811e8aa2d1661ecc7a408b27aaa3bdcaee408f52530c61a007953
	// PubComp: 03c6072aad509c88edad53756abc01f00f9e3feeb08fd748b4081964bae97e2532
	// Sin:     TfFabGCCATnbxfYRYfkCjM1RDC64aZfRJAu
}
