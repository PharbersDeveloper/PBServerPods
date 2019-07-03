// Package Model SandBoxIndex
package Model

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type SandBoxIndex struct {
	ID					string			`json:"-"`
	Id_					bson.ObjectId	`json:"-" bson:"_id"`
	AccountID			string			`json:"-" bson:"account-id"`
	FileMetaDataIDs		[]string		`json:"-" bson:"file-meta-data-ids"`
	FileMetaDatum		[]*FileMetaData	`json:"-"`
}


// GetID to satisfy jsonapi.MarshalIdentifier interface
func (f SandBoxIndex) GetID() string {
	return f.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (f *SandBoxIndex) SetID(id string) error {
	f.ID = id
	return nil
}

// GetReferences 设置关联关系详情参考JSONAPI定义的Reference
// https://jsonapi.org/format/#document-resource-object-relationships
func (f SandBoxIndex) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "fileMetaDatum",
			Name: "fileMetaDatum",
		},
	}
}

// GetReferencedIDs 获取关联ID
func (f SandBoxIndex) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range f.FileMetaDataIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "fileMetaDatum",
			Name: "fileMetaDatum",
		})
	}

	return result
}

// GetReferencedStructs 设置Reference内容
func (f SandBoxIndex) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range f.FileMetaDataIDs {
		result = append(result, f.FileMetaDatum[key])
	}

	return result
}

// SetToManyReferenceIDs 设置关联ID
func (f *SandBoxIndex) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "fileMetaDatum" {
		f.FileMetaDataIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs
func (f *SandBoxIndex) AddToManyIDs(name string, IDs []string) error {
	if name == "fileMetaDatum" {
		f.FileMetaDataIDs = append(f.FileMetaDataIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// GetConditionsBsonM MongoDB的拼接Condition
func (f *SandBoxIndex) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		}
	}
	return rst
}