package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/BreadTech/pki/internal/store"
	"github.com/BreadTech/pki/pkg/models"
)

func main() {
	db, err := gorm.Open("sqlite3", "pki.db")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	defer db.Close()

	if err = store.RunMigrations(db); err != nil {
		panic(err)
	}

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	keyBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		panic(err)
	}

	if err = db.Create(&models.Entity{
		ID:   uuid.Must(uuid.NewV4()),
		Name: "root-private-key",
		Kind: models.EntityTypePEM,
		Dat:  keyBytes,
	}).Error; err != nil {
		panic(err)
	}

	pubKey := privKey.Public()

	if keyBytes, err = x509.MarshalPKIXPublicKey(pubKey); err != nil {
		panic(err)
	}

	if err = db.Create(&models.Entity{
		ID:   uuid.Must(uuid.NewV4()),
		Name: "root-public-key",
		Kind: models.EntityTypePEM,
		Dat:  keyBytes,
	}).Error; err != nil {
		panic(err)
	}

	keys := []*models.Entity{}
	db.Find(&keys)

	for _, key := range keys {
		fmt.Println(key.Name, base64.StdEncoding.EncodeToString(key.Dat))
	}
}
