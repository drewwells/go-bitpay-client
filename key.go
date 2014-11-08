package go_bitpay_client

import (
	"fmt"
	"math/big"

	btcaddr "github.com/drewwells/gimme_bitcoin_address"
)

// Keygen generates a new private/public key pair.  These are used for pairing
// to the Bitpay API.
func Keygen() ([]byte, []byte, error) {

	p, b, err := btcaddr.Bitcoin_GenerateKeypair()
	_ = b
	fmt.Printf("Bitcoin len: %d\n", len(p.D.Bytes()))
	//privstr := gimme.Bitcoin_Prikey2WIF(p)

	fmt.Printf("int  %d %x\n", len(p.D.Bytes()), p.D.Bytes()) //THIS IS IT FUCK THESE ENCODING SCHEMAS
	//fmt.Printf("% #v\n", pretty.Formatter(b))
	pub := Public(b.X, b.Y)
	fmt.Printf("pub  %d %x\n", len(pub), pub)
	//fmt.Printf("pub %d %x\n", len(b.R.Bytes()), b.R.Bytes())

	//RIPEMD160 encode the int
	/*ripe := ripemd160.New()
	privhash := ripe.Sum(p.D.Bytes())
	fmt.Printf("hash %d %x\n", len(privhash), privhash)

	sha := sha256.New()
	shahash := sha.Sum(p.D.Bytes())
	fmt.Printf("sha  %d %x\n", len(shahash), shahash)

	md := md5.New()
	mdhash := md.Sum(p.D.Bytes())
	fmt.Printf("md5  %d %x\n", len(mdhash), mdhash)

	hx := make([]byte, hex.EncodedLen(len(p.D.Bytes())))
	hex.Encode(hx, p.D.Bytes())
	fmt.Printf("hex  %d %x\n", len(hx), hx)

	b58 := btcutil.Base58Encode(p.D.Bytes())
	fmt.Printf("b58  %d %s\n", len(b58), b58)

	ex := "2osLAeuhKKwu61eu7MnXpbGU7Rkb6j155aXo515mWoS91nDPRR5rJGgvG3VRGpPpGWo4AEU3HmqtsQUcuPm8aBancYp5kC81gWpY7PCPa7cFZg"
	fmt.Printf("ex  %d %s\n", len(ex), ex)*/

	return []byte{}, []byte{}, err
}

func Public(X, Y *big.Int) []byte {

	//sha256_h := sha256.New()
	/* Create a new RIPEMD160 Context */
	//ripemd160_h := ripemd160.New()

	/* Convert the public key to a byte sequence */
	pubkey_bytes := append(X.Bytes(), Y.Bytes()...)

	/* 1. Prepend 0x04 */
	return append([]byte{0x04}, pubkey_bytes...)

}
