CREATE TABLE IF NOT EXISTS "entities" (
	"id" UUID PRIMARY KEY,
	"name" TEXT UNIQUE,
	"kind" INTEGER,
	"dat" BLOB
);

CREATE TABLE IF NOT EXISTS "relationships" (
	"id" UUID,
	"entity_id" UUID,
	"name" TEXT,
	PRIMARY KEY ("id", "entity_id"),
	FOREIGN KEY ("entity_id") REFERENCES "entities" ("id") 
		ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS "events" (
	"id" UUID,
	"entity_id" UUID,
	"relationship_id" UUID,
	"metadata_id" UUID, 
	"action" TEXT,
	"timestamp" DATETIME,
	"signer" UUID,
	"signature" BLOB
);

CREATE TABLE IF NOT EXISTS "metadata" (
	"id" UUID,
	"key" TEXT,
	"value" BLOB,
	PRIMARY KEY ("id", "key")
);
