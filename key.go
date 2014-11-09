package go_bitpay_client

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	btcaddr "github.com/drewwells/gimme_bitcoin_address"
)

// Keygen generates a new private/public key pair using Bitcoin scheme
func Keygen() ([]byte, []byte, error) {

	p, b, err := btcaddr.Bitcoin_GenerateKeypair()
	//fmt.Printf("Bitcoin len: %d\n", len(p.D.Bytes()))
	//privstr := gimme.Bitcoin_Prikey2WIF(p)
	priv := p.D.Bytes()
	//fmt.Printf("int  %d %x\n", len(p.D.Bytes()), p.D.Bytes()) //THIS IS IT
	//fmt.Printf("% #v\n", pretty.Formatter(b))
	pub := public(b.X, b.Y)
	//fmt.Printf("pub  %d %x\n", len(pub), pub)

	return priv, pub, err
}

type PublicKey []byte

const (
	MAIN      = 0x01
	TESTNET   = 0x11
	EPHEMERAL = 0x02
)

// Sin accepts as input PublicKey and returns a SIN (Secure Identity Number)
// Learn more: https://en.bitcoin.it/wiki/Identity_protocol_v1
func Sin(key PublicKey) []byte {
	s := (sha256.New()).Sum(key)
	// Type1 persistant keys prefix with main 0x01, testnet 0x11
	// Type2 ephemeral 0x02
	bt := 0x02
	fmt.Println(Sin)
	_, _ = s, bt
	return []byte{}
}

// Accepts the X, Y point and returns the Public Key
func public(X, Y *big.Int) []byte {

	//sha256_h := sha256.New()
	/* Create a new RIPEMD160 Context */
	//ripemd160_h := ripemd160.New()

	/* Convert the public key to a byte sequence */
	pubkey_bytes := append(X.Bytes(), Y.Bytes()...)

	/* 1. Prepend 0x04 */
	return append([]byte{0x04}, pubkey_bytes...)

}

func PublicFromPrivate(priv []byte) []byte {
	d := new(big.Int).SetBytes(priv)
	curve := getCurve()
	Q := curve.Point_scalar_multiply(d, curve.G)
	//fmt.Printf("% #v\n", pretty.Formatter(Q))
	return public(Q.X, Q.Y)
}

func getCurve() *btcaddr.EllipticCurve {
	/* secp256k1 elliptic curve parameters */
	var curve = &btcaddr.EllipticCurve{}
	curve.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	curve.A, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000000", 16)
	curve.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	curve.G.X, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	curve.G.Y, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
	curve.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	curve.H, _ = new(big.Int).SetString("01", 16)

	return curve
}
