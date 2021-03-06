## go-bitpay-client

[![GoDoc](https://godoc.org/github.com/drewwells/go-bitpay-client?status.svg)](http://godoc.org/github.com/drewwells/go-bitpay-client)

Wrapper for the [Bitpay API](https://test.bitpay.com/api).  This library leverages the great [btcec](http://github.com/conformal/btcec) library for private/public authentication.

go-bitpay-client is still alpha and not all operations are currently supported.  The cli is installed via:

    go get github.com/drewwells/go-bitpay-client/go-bitpay
    go-bitpay

### Getting started

Use the node wrapper to generate your client tokens.  This part isn't implement in Go yet [see here](https://www.npmjs.org/package/bitpay).  Be sure your account is the correct type (test account for test.bitpay.com or prod account for www.bitpay.com).  For the rest of this setup, the docs will assume you are using the test environment.

    bitpay keygen
    bitpay pair

bitpay generates key files in the ~/.bitpay folder.  The keys generated now need to be hex encoded, this function will do that: https://github.com/bitpay/bitauth/blob/master/lib/bitauth.js#L12.  The results should be a long series of numbers and *LOWER CASE* letters.  Save this as your pub and priv values in the .env file.

Go back to go-bitpay-client project root and copy .env.sample to .env.  This file stores all the credentials needed for negotiating with the bitpay server.  Generate a token based off your *PUBLIC* key, this can be done a few different ways.  Try this one.  If it doesn't work please pull request to fix it, thanks!  In order to do this, you need to set the publickey value in your .env file.  This is your public key, hint: your public key the format is camelcase with numbers.

    go get github.com/drewwells/go-bitpay-client/go-bitpay
    go-bitpay token

You will receive a new token and a pairing code.  Login to your account and enter this pairing code on the [api-tokens](https://test.bitpay.com/api-tokens) page.  Add this token to the .env file as 'token'.

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

### Help

If you're getting stuck with invalid sin or invalid token, [bitauth](https://www.npmjs.org/package/bitauth) is the source of truth.  Try playing around with those examples with your credentials (located in ~/.bitpay).
