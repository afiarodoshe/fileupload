package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	FileID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Fid      string             `json:"fId" bson:"fId"`
	File     []byte             `json:"file" bson:"file"`
	FileName string             `json:"fileName" bson:"fileName"`
}
