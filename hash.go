package go_bitpay_client

import (
	"crypto/sha256"

	"code.google.com/p/go.crypto/ripemd160"
)

// DoubleShaHash performs 2 sha256 checksums and sets the bytes of dst.
func DoubleShaHash(out, b []byte) {
	sum := sha256.New()
	sum.Write(b)
	fst := sum.Sum(nil)
	sum.Reset()
	sum.Write(fst)
	copy(out, sum.Sum(nil))
}

// RimpHash performs a ripemd160 hash following a sha256 hash
// returning the ripemd160 checksum.
func RimpHash(out, in []byte) {
	sum := sha256.New()
	sum.Write(in)
	rim := ripemd160.New()
	rim.Write(sum.Sum(nil))
	copy(out, rim.Sum(nil))
}
