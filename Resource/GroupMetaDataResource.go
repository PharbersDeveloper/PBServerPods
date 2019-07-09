// Package Resource
package Resource

import (
	"SandBox/DataStorage"
	"SandBox/Model"
	"errors"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

// GroupMetaDataResource
type GroupMetaDataResource struct {
	GroupMetaDataStorage   *DataStorage.GroupMetaDataStorage
}

// NewGroupMetaDataResource Initialize Parameter And injection Storage Or Resource
func (s GroupMetaDataResource) NewGroupMetaDataResource(args []BmDataStorage.BmStorage) *GroupMetaDataResource {
	var dcs *DataStorage.GroupMetaDataStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "GroupMetaDataStorage" {
			dcs = arg.(*DataStorage.GroupMetaDataStorage)
		}
	}
	return &GroupMetaDataResource{
		GroupMetaDataStorage:    	dcs,
	}
}

func (s GroupMetaDataResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := s.GroupMetaDataStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s GroupMetaDataResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []Model.GroupMetaData
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.GroupMetaDataStorage.GetAll(r, int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.GroupMetaDataStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := Model.GroupMetaData{}
	count := s.GroupMetaDataStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s GroupMetaDataResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.GroupMetaDataStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s GroupMetaDataResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(Model.GroupMetaData)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid_Instance_Given"), "Invalid Instance Given", http.StatusBadRequest)
	}

	id := s.GroupMetaDataStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s GroupMetaDataResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.GroupMetaDataStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s GroupMetaDataResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(Model.GroupMetaData)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid_Instance_Given"), "Invalid Instance Given", http.StatusBadRequest)
	}

	err := s.GroupMetaDataStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
