# NKEYS

[![License Apache 2](https://img.shields.io/badge/License-Apache2-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![ReportCard](http://goreportcard.com/badge/nats-io/nkeys)](http://goreportcard.com/report/nats-io/nkeys)
[![Build Status](https://travis-ci.com/nats-io/nkeys.svg?branch=master)](http://travis-ci.com/nats-io/nkeys)
[![GoDoc](http://godoc.org/github.com/nats-io/nkeys?status.svg)](http://godoc.org/github.com/nats-io/nkeys)
[![Coverage Status](https://coveralls.io/repos/github/nats-io/nkeys/badge.svg?branch=master&service=github)](https://coveralls.io/github/nats-io/nkeys?branch=master)

A public-key signature system based on [Secp256k1](https://pkg.go.dev/github.com/ethereum/go-ethereum/crypto@v1.10.5) for the NATS ecosystem.

## About

The NATS ecosystem will be moving to [Secp256k1](https://pkg.go.dev/github.com/ethereum/go-ethereum/crypto@v1.10.5) keys for identity, authentication and authorization for entities such as Accounts, Users, Servers and Clusters.

`TODO: update this section about Secp256k1`

## Installation

Use the `go` command:

	$ go get github.com/nats-io/nkeys

## nk - Command Line Utility

Located under the nk [directory](https://github.com/nats-io/nkeys/tree/master/nk).

## Basic API Usage
```go

// Create a new User KeyPair
user, _ := nkeys.CreateUser()

// Sign some data with a full key pair user.
data := []byte("Hello World")
sig, _ := user.Sign(data)

// Verify the signature.
err = user.Verify(data, sig)

// Access the seed, the only thing that needs to be stored and kept safe.
// seed = "SUAKYRHVIOREXV7EUZTBHUHL7NUMHPMAS7QMDU3GTIUWEI5LDNOXD43IZY"
seed, _ := user.Seed()

// Access the public key which can be shared.
// publicKey = "UD466L6EBCM3YY5HEGHJANNTN4LSKTSUXTH7RILHCKEQMQHTBNLHJJXT"
publicKey, _ := user.PublicKey()

// Create a full User who can sign and verify from a private seed.
user, _ = nkeys.FromSeed(seed)

// Create a User who can only verify signatures via a public key.
user, _ = nkeys.FromPublicKey(publicKey)

// Create a User KeyPair with our own random data.
var rawSeed [40]byte
_, err := io.ReadFull(rand.Reader, rawSeed[:])  // Or some other random source.
user2, _ := nkeys.FromRawSeed(PrefixByteUser, rawSeed)

```

## License

Unless otherwise noted, the NATS source files are distributed
under the Apache Version 2.0 license found in the LICENSE file.

