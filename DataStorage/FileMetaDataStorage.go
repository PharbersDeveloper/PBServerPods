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

// FileMetaDataStorage 注入MongoDB
type FileMetaDataStorage struct {
	db *BmMongodb.BmMongodb
}

// NewFileMetaDataStorage initialize parameter
func (s FileMetaDataStorage) NewFileMetaDataStorage(args []BmDaemons.BmDaemon) *FileMetaDataStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &FileMetaDataStorage{mdb}
}

// GetAll of the model
func (s FileMetaDataStorage) GetAll(r api2go.Request, skip int, take int) []*Model.FileMetaData {
	in := Model.FileMetaData{}
	var out []*Model.FileMetaData
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
func (s FileMetaDataStorage) GetOne(id string) (Model.FileMetaData, error) {
	in := Model.FileMetaData{ID: id}
	out := Model.FileMetaData{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("FileMetaData for id %s not found", id)
	return Model.FileMetaData{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *FileMetaDataStorage) Insert(c Model.FileMetaData) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *FileMetaDataStorage) Delete(id string) error {
	in := Model.FileMetaData{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("FileMetaData with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing model
func (s *FileMetaDataStorage) Update(c Model.FileMetaData) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("FileMetaData with id does not exist")
	}

	return nil
}

// Count MongoDB Query amount
func (s *FileMetaDataStorage) Count(req api2go.Request, c Model.FileMetaData) int {
	r, _ := s.db.Count(req, &c)
	return r
}