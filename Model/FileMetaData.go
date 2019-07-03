// Package Model FileMetaData
package Model

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// FileMetaData 文件源数据
type FileMetaData struct {
	ID				string			`json:"-"`
	Id_ 			bson.ObjectId	`json:"-" bson:"_id"`
	Name			string			`json:"name" bson:"name"`
	Extension		string			`json:"extension" bson:"extension"`
	Created 		int64			`json:"created" bson:"created"`
	Kind 			string			`json:"kind" bson:"kind"`
	OwnerID			string			`json:"-" bson:"owner-id"`
	GroupID			string			`json:"-" bson:"group-id"`
	Mod 			string			`json:"mod" bson:"mod"`
	FileVersionIDs	[]string		`json:"-" bson:"file-version-ids"`
	FileVersions	[]*FileVersion	`json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (f FileMetaData) GetID() string {
	return f.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (f *FileMetaData) SetID(id string) error {
	f.ID = id
	return nil
}

// GetReferences 设置关联关系详情参考JSONAPI定义的Reference
// https://jsonapi.org/format/#document-resource-object-relationships
func (f FileMetaData) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "fileVersions",
			Name: "fileVersions",
		},
	}
}

// GetReferencedIDs 获取关联ID
func (f FileMetaData) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range f.FileVersionIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "fileVersions",
			Name: "fileVersions",
		})
	}
	
	return result
}

// GetReferencedStructs 设置Reference内容
func (f FileMetaData) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range f.FileVersionIDs {
		result = append(result, f.FileVersions[key])
	}
	
	return result
}

// SetToManyReferenceIDs 设置关联ID
func (f *FileMetaData) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "fileVersions" {
		f.FileVersionIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs
func (f *FileMetaData) AddToManyIDs(name string, IDs []string) error {
	if name == "fileVersions" {
		f.FileVersionIDs = append(f.FileVersionIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// GetConditionsBsonM MongoDB的拼接Condition
func (f *FileMetaData) GetConditionsBsonM(parameters map[string][]string) bson.M {
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