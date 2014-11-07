package go_bitpay_client

import (
	"crypto/sha256"

	"github.com/conformal/btcutil"
)

const BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// DoubleSha256 wraps calling Sha256 twice
func DoubleSha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	sum := hasher.Sum(nil)
	hasher.Reset()
	hasher.Write(sum)
	return hasher.Sum(nil)
}

// Base58Check accepts base10 byte string and returns
// base58 encoded verification string as defined:
// https://en.bitcoin.it/wiki/Base58Check_encoding#Creating_a_Base58Check_string
func Base58Check(h []byte) string {
	check := make([]byte, len(h)+4)
	copy(check, h)
	hash := DoubleSha256(h)[:4]

	for i := 0; i < 4; i++ {
		check[i+len(h)] = hash[i]
	}
	return btcutil.Base58Encode(check)
}
