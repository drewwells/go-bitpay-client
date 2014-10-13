package go_bitpay_client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"

	"code.google.com/p/gcfg"
)

type Config struct {
	Global struct {
		Key, End string
	}
}

var cfg Config

func init() {
	err := gcfg.ReadFileInto(&cfg, ".env")
	if err != nil {
		panic(err)
	}
	if cfg.Global.Key == "" {
		log.Fatal("Please register an API key at bitpay.com")
	}
}

// Invoice creates a currency invoice
// curl https://bitpay.com/api/invoice \
//   -u ApiKey \
//   -d price=10.00 \
//   -d currency=USD
func Invoice() {
	stringResponse("/api/invoice", url.Values{
		"price":    {"10.00"},
		"currency": {"USD"},
	})
}

func stringResponse(path string, data url.Values) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	req, err := http.NewRequest("POST",
		cfg.Global.End+path, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.ParseForm()
	req.SetBasicAuth(cfg.Global.Key, "")
	req.Header.Add("User-Agent", "curl/7.30.0") //Bitpay/v1 GoBindings/1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	fmt.Println(req.URL)
	res, err := client.Do(req)
	fmt.Printf("% #v\n", req)
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
