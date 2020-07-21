module github.com/BreadTech/pki

go 1.13

replace breadtech/interface => ./pkg/bi/go

require (
	breadtech/interface v0.0.1
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/jinzhu/gorm v1.9.15
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
)
