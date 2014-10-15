package go_bitpay_client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"crypto/sha256"

	"code.google.com/p/gcfg"
	"github.com/conformal/btcec"
)

type Config struct {
	Global struct {
		Id, Pub, Priv, End, Token string
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

func Token() []byte {
	resp := stringResponse("/tokens", //?nonce="+
		//strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
		"POST", url.Values{
			"guid":  {"1"},
			"label": {"node-bitpay-client-dwells-mac2"},
			"id":    {cfg.Global.Id},
		}, true)
	return resp
}

// Invoice creates a currency invoice
// curl https://bitpay.com/api/invoice \
//   -u ApiKey \
//   -d price=10.00 \
//   -d currency=USD
func Invoice() []byte {

	return stringResponse("/invoices", "POST", url.Values{
		"price":    {"10.00"},
		"currency": {"USD"},
		//"nonce":    {strconv.FormatInt(time.Now().UnixNano()/1000000, 10)},
		"guid":  {"553c5ca8-b3e6-c9b2-8a29-e203ccd9d45g"},
		"token": {cfg.Global.Token},
	}, false)
}

func Rates() []byte {
	return stringResponse("/rates", "GET", url.Values{}, true)
}

func stringResponse(path, method string, data url.Values, public bool) []byte {
	unwrap := make(map[string]interface{})

	for i := range data {
		if i == "nonce" {
			i64, _ := strconv.ParseInt(data[i][0], 10, 64)
			unwrap[i] = i64
		} else if i == "price" {
			f64, _ := strconv.ParseFloat(data[i][0], 64)
			unwrap[i] = f64
		} else {
			unwrap[i] = data[i][0]
		}
	}
	var bs []byte
	if len(unwrap) > 0 {
		bs, _ = json.Marshal(unwrap)
	}
	var body io.Reader
	if len(bs) > 0 {
		body = strings.NewReader(string(bs))
	}
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	path = cfg.Global.End + path
	req, err := http.NewRequest(method,
		path,
		body)
	if err != nil {
		log.Fatal(err)
	}
	pub := cfg.Global.Pub
	contract := path + string(bs)
	sign := signMessage(cfg.Global.Priv, contract)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-accept-version", "2.0.0")
	if !public {
		req.Header.Add("x-identity", pub)
		req.Header.Add("x-signature", sign)
	}
	req.Header.Add("Content-Length",
		strconv.FormatInt(req.ContentLength, 10))
	// Dump the request before client eats the headers
	// rawr, _ := httputil.DumpRequestOut(req, true)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Raw dump of request/response
	// fmt.Println("request", string(rawr))
	// dump, _ := httputil.DumpResponse(res, true)
	// fmt.Println("response", string(dump))
	res.Body.Close()
	return bytes
}

func signMessage(key, message string) string {
	log.Print("signing", message)
	// Decode a hex-encoded private key.
	pkBytes, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)
	_ = pubKey
	// Sign a message using the private key.

	h := sha256.New()
	io.WriteString(h, message)
	signhash := h.Sum(nil)
	// Do not use this, sha256 hash is used for token creation
	// messageHash := btcwire.DoubleSha256([]byte(message))
	signature, err := privKey.Sign(signhash)
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
	// fmt.Printf("Signature Verified against pubkey? %v\n", verified)
	return fmt.Sprintf("%x", signature.Serialize())
}
