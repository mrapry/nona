package library

import (
	"reflect"
	"strings"

	"github.com/Kamva/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// ToBSON function set struct to bson M
func ToBSON(structure interface{}) bson.M {
	t := reflect.TypeOf(structure)
	v := reflect.ValueOf(structure)

	// set bson encoding
	bsonEncoding := []bson.M{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags := field.Tag.Get("bson")

		// set invalid tag
		invalidTag := []string{"", "-", ",inline"}
		isInvalidTag := false

		// validate tag
		for _, untag := range invalidTag {
			if tags == untag {
				isInvalidTag = true
			}
		}
		if isInvalidTag {
			continue
		}

		// set bson name
		tag := strings.Split(tags, ",")
		bsonName := tag[0]

		// validate value
		vField := v.Field(i)
		if !reflect.DeepEqual(vField.Interface(), reflect.Zero(vField.Type()).Interface()) {
			bsonEncoding = append(bsonEncoding, bson.M{bsonName: vField.Interface()})
		}
	}

	return bson.M{operator.And: bsonEncoding}
}
