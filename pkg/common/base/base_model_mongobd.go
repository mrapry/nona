package base

import "time"

//ModelMongo structure define model for MongoDB
type ModelMongo struct {
	ID string             `bson:"_id"`
	IsDeleted bool        `bson:"is_deleted"`
	CreatedBy string      `bson:"created_by"`
	ModifiedBy string     `bson:"modified_by"`
	CreatedAt time.Time   `bson:"created_at"`
	ModifiedAt *time.Time `bson:"modified_at"`
}