package go_bitpay_client

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"

	"code.google.com/p/gcfg"
	"github.com/conformal/btcec"
	"github.com/conformal/btcwire"
)

type Config struct {
	Global struct {
		Pub, Priv, End string
	}
}

var cfg Config

func init() {
	err := gcfg.ReadFileInto(&cfg, ".env")
	if err != nil {
		panic(err)
	}
	if cfg.Global.Pub == "" {
		log.Fatal("Please register an API key at bitpay.com and sign it with bitpay cli")
	}
}

// Invoice creates a currency invoice
// curl https://bitpay.com/api/invoice \
//   -u ApiKey \
//   -d price=10.00 \
//   -d currency=USD
func Invoice() {

	stringResponse("/tokens", url.Values{
	//"price":    {"10.00"},
	//"currency": {"USD"},
	})
}

func stringResponse(path string, data url.Values) {
	unwrap := make(map[string]string)

	for i := range data {
		unwrap[i] = data[i][0]
	}

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	path = cfg.Global.End + path + "?nonce=" + strconv.FormatInt(time.Now().Unix(), 10)
	req, err := http.NewRequest("GET",
		path,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	req.ParseForm()
	pub := cfg.Global.Pub
	sign := signMessage(cfg.Global.Priv, path)
	req.Header.Add("x-pubkey", pub)
	req.Header.Add("x-accept-version", "2.0.0")
	req.Header.Add("x-signature", sign)
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(req)
	fmt.Printf("% #v\n", req.Header)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	log.Printf("%s\n", string(bytes))
}

func signMessage(key, message string) string {
	fmt.Println("signing", message)
	// Decode a hex-encoded private key.
	pkBytes, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)
	_ = pubKey
	// Sign a message using the private key.
	messageHash := btcwire.DoubleSha256([]byte(message))
	signature, err := privKey.Sign(messageHash)
	if err != nil {
		log.Fatal(err)
	}

	// Serialize and display the signature.
	//
	// NOTE: This is commented out for the example since the signature
	// produced uses random numbers and therefore will always be different.
	// fmt.Printf("Serialized Signature: %x\n", signature.Serialize())

	// Verify the signature for the message using the public key.
	// verified := signature.Verify(messageHash, pubKey)
	// fmt.Printf("Signature Verified? %v\n", verified)
	return fmt.Sprintf("%x", signature.Serialize())
}
