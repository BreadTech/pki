# pki
Public Key Infrastructure tools

## Introduction
The [OpenSSL](https://www.openssl.org/docs/manmaster/man1/) tool is the de-facto standard for all your cryptography needs. However, after years of appending new subcommands to the tool without the appropriate refactoring, it has become a monolith of many features beyond the simple pki routines of public/private key generation as well as data signing and verification. `pkictl` is the command-line tool in this repository that provides a simple and clean interface to those methods.

## Getting Started

### Prerequisites
[Go 1.13](https://golang.org/dl/)

### To compile
```
git clone https://github.com/BreadTech/pki
cd pki
go build ./cmd/pkictl
```

### Generating a key

#### [RSA](https://en.wikipedia.org/wiki/RSA_%28cryptosystem%29)
```
./pkictl generate rsa
```
will generate a 2048-bit RSA private key (TODO: parameterize bit size)

#### [Elliptic-Curve](https://en.wikipedia.org/wiki/Elliptic-curve_cryptography)
```
./pkictl generate ecc
```
will generate an elliptic-curve cryptography private key using the P224 curve (TODO: parameterize curve)

#### [ED25519](https://en.wikipedia.org/wiki/EdDSA)
```
./pkictl generate edc
```
will generate an Edwards-curve private key.
