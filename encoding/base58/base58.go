/*
Base58 is a group of binary-to-text encoding schemes used to represent large integers as alphanumeric text. It is similar to Base64 but has been modified to avoid both non-alphanumeric characters and letters which might look ambiguous when printed. It is therefore designed for human users who manually enter the data, copying from some visual source, but also allows easy copy and paste because a double-click will usually select the whole string.

Compared to Base64, the following letters have been omitted from the alphabet: 0 (zero), O (capital o), I (capital i) and l (lower case L) as well as the non-alphanumeric characters + (plus) and / (slash). In contrast to Base64, the digits of the encoding don't line up well with byte boundaries of the original data. For this reason, the method is well-suited to encode large integers, but not designed to encode longer portions of binary data. The actual order of letters in the alphabet depends on the application, which is the reason why the term “Base58” alone is not enough to fully describe the format.

base58 returns encoded text suitable for use with Bitcoin.  Bitcoin compatible base58 does not pad like fixed width base58.  As a result, the size of the returned slice can be different for the same sized input.
*/
package base58

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

const base58table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func EncodedLen() {}

func DecodedLen() {}

var ErrInvalid = errors.New("encoding/base58: invalid character found")

const radix = 58

// Decode decodes src base10 string and returns the base58 encoded string
// and size of the result.
func Decode(src []byte) ([]byte, int, error) {
	b := string(src)
	answer := big.NewInt(0)
	j := big.NewInt(1)

	for i := len(b) - 1; i >= 0; i-- {
		tmp := strings.IndexAny(base58table, string(b[i]))
		if tmp == -1 {
			fmt.Println(b)
			return []byte(""), 0,
				errors.New("encoding/base58: invalid character found: ~" +
					string(b[i]) + "~")
		}
		idx := big.NewInt(int64(tmp))
		tmp1 := big.NewInt(0)
		tmp1.Mul(j, idx)

		answer.Add(answer, tmp1)
		j.Mul(j, big.NewInt(radix))
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != base58table[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen, flen)
	copy(val[numZeros:], tmpval)
	return val, len(val), nil
}

// DecodeString returns the bytes represented by the base58 string s
func DecodeString(s string) ([]byte, error) {
	dst, _, err := Decode([]byte(s))
	return dst, err
}

// DoubleSha256 wraps calling Sha256 twice
func doubleSha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	sum := hasher.Sum(nil)
	hasher.Reset()
	hasher.Write(sum)
	return hasher.Sum(nil)
}

// Base58Check accepts a byte string and returns base58 encoded
// verification string as defined:
// https://en.bitcoin.it/wiki/Base58Check_encoding#Creating_a_Base58Check_string
func Check(h []byte) string {
	check := make([]byte, len(h)+4)
	copy(check, h)
	hash := doubleSha256(h)[:4]

	for i := 0; i < 4; i++ {
		check[i+len(h)] = hash[i]
	}
	dst, _ := Encode(check)
	return string(dst)
}

// The radix of the base58 encoding system.
const Radix = len(base58table)

// Bits of entropy per base 58 digit.
var BitsPerDigit = math.Log2(float64(Radix))

// Base58MaxEncodedLen returns the maximum possible length
// of a base58 encoding.  This number may be larger than the
// encoded slice.
func MaxEncodedLen(b []byte) int {
	maxlen := int(math.Ceil(float64(len(b)) / BitsPerDigit * 8))
	return maxlen
}

// Encode creates Bitcoin compatible Base58 encoded strings
// from a byte slice.  The length is variable based on same
// sized input slice.
func Encode(src []byte) ([]byte, int) {

	var dst []byte
	x := new(big.Int).SetBytes(src)
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s := ""

	/* While x > 0 */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(base58table[r.Int64()]) + s
		dst = append(dst, base58table[r.Int64()])
	}

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range src {
		if v != 0 {
			break
		}
		dst = append(dst, base58table[0])
	}

	for i := 0; i < len(dst)/2; i++ {
		dst[i], dst[len(dst)-1-i] =
			dst[len(dst)-1-i], dst[i]
	}
	return dst, len(dst)
}

// EncodeToString returns a string from a byte slice.
func EncodeToString(src []byte) string {
	dst, _ := Encode(src)
	return string(dst)
}
