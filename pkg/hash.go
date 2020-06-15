package pkg

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

var (
	hashes = map[string]crypto.Hash{
		"sha224": crypto.SHA224,
		"sha256": crypto.SHA256,
		"sha384": crypto.SHA384,
		"sha512": crypto.SHA512,
	}
	hashFuncs = map[crypto.Hash]func() hash.Hash{
		crypto.SHA224: sha256.New224,
		crypto.SHA256: sha256.New,
		crypto.SHA384: sha512.New384,
		crypto.SHA512: sha512.New,
	}
)

func Hashes() []string {
	hashList := make([]string, len(hashes))
	i := 0
	for k := range hashes {
		hashList[i] = k
		i++
	}
	return hashList
}

func IsValidHash(hashType string) bool {
	_, exists := hashes[hashType]
	return exists
}

func Hash(hashType string, dat []byte) (crypto.Hash, []byte) {
	hashEnum, ok := hashes[hashType]
	if !ok {
		panic("Invalid hash type")
	}
	hasher := hashFuncs[hashEnum]()
	hasher.Write(dat)
	return hashEnum, hasher.Sum(nil)
}

func PSSOption(hashType string) *rsa.PSSOptions {
	return &rsa.PSSOptions{
		SaltLength: 8,
		Hash:       hashes[hashType],
	}
}
