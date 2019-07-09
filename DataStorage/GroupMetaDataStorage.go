// Package DataStorage
package DataStorage

import (
	"SandBox/Model"
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"net/http"
)

// GroupMetaDataStorage 注入MongoDB
type GroupMetaDataStorage struct {
	db *BmMongodb.BmMongodb
}

// NewGroupMetaDataStorage initialize parameter
func (s GroupMetaDataStorage) NewGroupMetaDataStorage(args []BmDaemons.BmDaemon) *GroupMetaDataStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &GroupMetaDataStorage{mdb}
}

// GetAll of the model
func (s GroupMetaDataStorage) GetAll(r api2go.Request, skip int, take int) []*Model.GroupMetaData {
	in := Model.GroupMetaData{}
	var out []*Model.GroupMetaData
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty model
func (s GroupMetaDataStorage) GetOne(id string) (Model.GroupMetaData, error) {
	in := Model.GroupMetaData{ID: id}
	out := Model.GroupMetaData{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("GroupMetaData for id %s not found", id)
	return Model.GroupMetaData{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *GroupMetaDataStorage) Insert(c Model.GroupMetaData) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *GroupMetaDataStorage) Delete(id string) error {
	in := Model.GroupMetaData{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("GroupMetaData with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing model
func (s *GroupMetaDataStorage) Update(c Model.GroupMetaData) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("GroupMetaData with id does not exist")
	}

	return nil
}

// Count MongoDB Query amount
func (s *GroupMetaDataStorage) Count(req api2go.Request, c Model.GroupMetaData) int {
	r, _ := s.db.Count(req, &c)
	return r
}