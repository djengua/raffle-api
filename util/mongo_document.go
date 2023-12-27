package util

import (
	"go.mongodb.org/mongo-driver/bson"
)

func StructToDocument(i interface{}) (doc *bson.D, err error) {

	data, err := bson.Marshal(i)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)

	return
}

// func StructToDocumentSave(i interface{}) (doc *bson.D, err error) {

// 	data, err := bson.Marshal(i)
// 	if err != nil {
// 		return
// 	}
// 	err = bson.Unmarshal(data, &doc)
// 	return
// }
