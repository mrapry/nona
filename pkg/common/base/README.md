#Common Base

##Base


This package use for define base model like  MongoDB and SQL created. So you no need create multiple field in same model in struct

Sample Code
```go
//ModelSQL structure define model for SQL
type ModelSQL struct {
	ID string             `json:"id"`
	IsDeleted bool        `json:"is_deleted"`
	CreatedBy string      `json:"created_by"`
	ModifiedBy string     `json:"modified_by"`
	CreatedAt time.Time   `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}

//ModelMongo structure define model for MongoDB
type ModelMongo struct {
	ID string             `bson:"_id"`
	IsDeleted bool        `bson:"is_deleted"`
	CreatedBy string      `bson:"created_by"`
	ModifiedBy string     `bson:"modified_by"`
	CreatedAt time.Time   `bson:"created_at"`
	ModifiedAt *time.Time `bson:"modified_at"`
}
```

For use this model you just type like this :

```go
Import "nona/pkg/common/base"


type Account struct{
	base.ModelMongo{} // like this
	Name string       `bson:"name"`
	Username string   `bson:"username"`
	Password string   `bson:"password"`
}
```