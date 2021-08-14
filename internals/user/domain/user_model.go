package domain

import "nona/pkg/common/base"

//collectionUser define name collection in db mongo
const collectionUser = "user"

//User structure
type User struct {
	base.ModelMongo
	Name     string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

// CollectionName for Configurations model
func (m *User) CollectionName() string {
	return collectionUser
}
