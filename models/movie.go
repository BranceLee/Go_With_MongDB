package models

import (
	"gopkg.in/mgo.v2/bson"
)

/*
** 定义数据库结构, bson => 对应数据库的key, json => 对应json 格式的key
*/
type Movies struct{
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name`
	CoverImage string `bson:"cover_image" json:"cover_image"`
	Description string `bson:"description" json:"description"`
}