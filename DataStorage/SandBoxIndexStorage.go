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

// SandBoxIndexStorage 注入MongoDB
type SandBoxIndexStorage struct {
	db *BmMongodb.BmMongodb
}

// NewSandBoxIndexStorage initialize parameter
func (s SandBoxIndexStorage) NewSandBoxIndexStorage(args []BmDaemons.BmDaemon) *SandBoxIndexStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &SandBoxIndexStorage{mdb}
}

// GetAll of the model
func (s SandBoxIndexStorage) GetAll(r api2go.Request, skip int, take int) []*Model.SandBoxIndex {
	in := Model.SandBoxIndex{}
	var out []*Model.SandBoxIndex
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
func (s SandBoxIndexStorage) GetOne(id string) (Model.SandBoxIndex, error) {
	in := Model.SandBoxIndex{ID: id}
	out := Model.SandBoxIndex{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("CitySalesReport for id %s not found", id)
	return Model.SandBoxIndex{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *SandBoxIndexStorage) Insert(c Model.SandBoxIndex) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *SandBoxIndexStorage) Delete(id string) error {
	in := Model.SandBoxIndex{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("CitySalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing model
func (s *SandBoxIndexStorage) Update(c Model.SandBoxIndex) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("CitySalesReport with id does not exist")
	}

	return nil
}

// Count MongoDB Query amount
func (s *SandBoxIndexStorage) Count(req api2go.Request, c Model.SandBoxIndex) int {
	r, _ := s.db.Count(req, &c)
	return r
}