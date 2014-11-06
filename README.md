## go-bitpay-client

Wrapper for the [Bitpay API](https://test.bitpay.com/api).  This library leverages the great [btcec](http://github.com/conformal/btcec) library for private/public authentication.

go-bitpay-client is still alpha and not all operations are currently supported.  The cli is installed via:

    go get github.com/drewwells/go-bitpay-client/bitpay
	bitpay

Generate a token based off your private key.  Get a private key by signing up for an account with [Bitpay](https://test.bitpay.com/start).


### Development

Getting started with Go https://golang.org/doc/install

Learning Go

* http://www.golang-book.com/
* Great book! [Programming Go](http://www.amazon.com/Programming-Go-Creating-Applications-Developers/dp/0321774639)
* https://gobyexample.com

Setup an .env file with your private, public keys and user id.  Copy .env.sample and fill in these values.  Keys are currently generated via debugging the Node API.  Fork or checkout this project.

	go get github.com/drewwells/go-bitpay-client
	cd $GOPATH/src/github.com/drewwells/go-bitpay-client
	go run bitpay/main.go
