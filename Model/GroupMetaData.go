// Package Model FileVersion
package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// GroupMetaData 文件版本
type GroupMetaData struct{
	ID			string			`json:"-"`
	Id_			bson.ObjectId	`json:"-" bson:"_id"`
	RoleID		string			`json:"role-id" bson:"role-id"`
	GroupID		string			`json:"group-id" bson:"group-id"`
	AccountID	string			`json:"account-id" bson:"account-id"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (f GroupMetaData) GetID() string {
	return f.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (f *GroupMetaData) SetID(id string) error {
	f.ID = id
	return nil
}

func (f *GroupMetaData) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "account-id":
			rst[k] = v[0]
		case "group-id":
			rst[k] = v[0]
		}
	}
	return rst
}