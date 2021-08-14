package base

import "time"

//ModelMongo structure define domain for MongoDB
type ModelMongo struct {
	ID         string     `bson:"_id" json:"id"`
	IsDeleted  bool       `bson:"is_deleted" json:"is_deleted"`
	CreatedBy  string     `bson:"created_by" json:"created_by"`
	ModifiedBy string     `bson:"modified_by" json:"modified_by"`
	CreatedAt  time.Time  `bson:"created_at" json:"created_at"`
	ModifiedAt *time.Time `bson:"modified_at" json:"modified_at"`
}
