package go_bitpay_client

import (
	"encoding/hex"
	"testing"
)

var sampleKey = []byte("02F840A04114081690223B7069071A70D6DABB891763B638CC20C7EC3BD58E6C86")

func TestShaHash(t *testing.T) {
	src, dst := sampleKey, make([]byte, 32)
	DoubleShaHash(dst, src)

	ehash := "eb5b348d7000a12b4de1ef7450b223e47ff7b2716f84e74a7c5816cc300e3300"
	hash := hex.EncodeToString(dst)
	if ehash != hash {
		t.Errorf("got:\n%s\nwanted:\n%s", hash, ehash)
	}
}

func TestRimpHash(t *testing.T) {
	dst := make([]byte, 20)
	src := make([]byte, hex.DecodedLen(len(sampleKey)))
	hex.Decode(src, sampleKey)
	RimpHash(dst, src)

	e := "cb1f4a4d793731842732c153b8e9923bdb462553"
	if e != hex.EncodeToString(dst) {
		t.Errorf("got:\n%x\nwanted:\n%s", string(dst), e)
	}
}

var dhash []byte

func BenchmarkDoubleShaHash(b *testing.B) {
	in := []byte("02F840A04114081690223B7069071A70D6DABB891763B638CC20C7EC3BD58E6C86")
	test := make([]byte, 32)
	for n := 0; n < b.N; n++ {
		DoubleShaHash(test, in)
	}
	dhash = test
}
