package go_bitpay_client

import (
	"math/big"

	btcaddr "github.com/drewwells/gimme_bitcoin_address"
	"github.com/drewwells/go-bitpay-client/encoding/base58"
)

type PublicKey []byte

const (
	MAIN      = 0x01
	TESTNET   = 0x11
	EPHEMERAL = 0x02
)

// Create a new Bitcoin compliant private and public key
// pair.
func GenerateKeyGen() ([]byte, []byte, error) {
	p, b, err := btcaddr.Bitcoin_GenerateKeypair()
	priv := p.D.Bytes()

	pub := public(b.X, b.Y, true)

	return priv, pub, err
}

// Sin accepts as input hex encoded compressed PublicKey and
// returns a SIN (Secure Identity Number)
// Learn more: https://en.bitcoin.it/wiki/Identity_protocol_v1#SIN_record
func Sin(key []byte) []byte {
	//rhash := btc.Rimp160AfterSha256(key)
	rhash := make([]byte, 20)
	RimpHash(rhash, key)
	// fmt.Printf("step2: %x\n", rhash)
	// Type1 persistant keys prefix with main 0x01, testnet 0x11
	// Type2 ephemeral 0x02
	bt := []byte{0x0F, 0x02}
	digest := make([]byte, 22)
	copy(digest, bt)
	for i := range rhash {
		digest[i+2] = rhash[i]
	}
	//fmt.Printf("step3: %x\n", digest)
	hash := base58.DoubleSha256(digest)
	//fmt.Printf("step4: %x\n", hash)

	checksum := hash[:4]
	//fmt.Printf("step5: %x\n", checksum)

	sin := append(digest, checksum...)
	//fmt.Printf("step6: %x\n", sin)
	encodedSin, _ := base58.Encode([]byte(sin))
	//fmt.Printf("step7: %s\n", encodedSin)

	return encodedSin
}

// Accepts the X, Y point and returns the Public Key
func public(X, Y *big.Int, compressed bool) []byte {
	var pubkey_bytes []byte
	if compressed {
		prefix := make([]byte, 1)
		mod := Y.Mod(Y, big.NewInt(2)).Cmp(big.NewInt(0))
		if mod == 0 {
			prefix[0] = 0x02
		} else {
			prefix[0] = 0x03
		}
		pubkey_bytes = append(prefix, X.Bytes()...)
		return pubkey_bytes
	}

	/* Convert the public key to a byte sequence */
	pubkey_bytes = append(X.Bytes(), Y.Bytes()...)

	/* 1. Prepend 0x04 */
	return append([]byte{0x04}, pubkey_bytes...)

}

// Reference for generating compressed keys
// github.com/piotrnar/gocoin/blob/master/lib/btc
func PublicFromPrivate(priv []byte, compressed bool) []byte {
	d := new(big.Int).SetBytes(priv)
	curve := getCurve()
	Q := curve.Point_scalar_multiply(d, curve.G)

	return public(Q.X, Q.Y, compressed)
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
