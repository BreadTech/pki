package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Entity struct {
	ID   uuid.UUID
	Name string
	Kind EntityType
	Dat  []byte
}

type Relationship struct {
	ID       uuid.UUID
	EntityID uuid.UUID
}

type Metadata struct {
	ID    uuid.UUID
	Key   string
	Value []byte
}

type Event struct {
	ID             uuid.UUID
	EntityID       uuid.UUID
	RelationshipID uuid.UUID
	MetadataID     uuid.UUID
	Action         string
	Timestamp      time.Time
	Signer         uuid.UUID
	Signature      []byte
}
