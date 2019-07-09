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

// FileMetaDataResource
type FileMetaDataResource struct {
	FileMetaDataStorage   	*DataStorage.FileMetaDataStorage
	SandBoxIndexStorage		*DataStorage.SandBoxIndexStorage
	GroupMetaDataStorage	*DataStorage.GroupMetaDataStorage
}

// NewFileMetaDataResource Initialize Parameter And injection Storage Or Resource
func (s FileMetaDataResource) NewFileMetaDataResource(args []BmDataStorage.BmStorage) *FileMetaDataResource {
	var dcs *DataStorage.FileMetaDataStorage
	var sbi *DataStorage.SandBoxIndexStorage
	var gmds *DataStorage.GroupMetaDataStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "FileMetaDataStorage" {
			dcs = arg.(*DataStorage.FileMetaDataStorage)
		} else if tp.Name() == "SandBoxIndexStorage" {
			sbi = arg.(*DataStorage.SandBoxIndexStorage)
		} else if tp.Name() == "GroupMetaDataStorage" {
			gmds = arg.(*DataStorage.GroupMetaDataStorage)
		}
	}
	return &FileMetaDataResource{
		FileMetaDataStorage: dcs,
		SandBoxIndexStorage: sbi,
		GroupMetaDataStorage: gmds,
	}
}

func (s FileMetaDataResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	sandBoxIndicesID, sok := r.QueryParams["sandBoxIndicesID"]

	if sok {
		modelRootID := sandBoxIndicesID[0]

		modelRoot, err := s.SandBoxIndexStorage.GetOne(modelRootID)

		if err != nil {
			return  &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.FileMetaDataIDs

		result := s.FileMetaDataStorage.GetAll(r, -1, -1)

		return &Response{Res: result}, nil
	}

	// TODO : 根据OAuth account id 查询GroupID与Role

	r.QueryParams["group-id"] = []string{"5cb9952d82a4a74375fa41fd"}

	result := s.FileMetaDataStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s FileMetaDataResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []Model.FileMetaData
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
		for _, iter := range s.FileMetaDataStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.FileMetaDataStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := Model.FileMetaData{}
	count := s.FileMetaDataStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s FileMetaDataResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.FileMetaDataStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s FileMetaDataResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(Model.FileMetaData)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid_Instance_Given"), "Invalid Instance Given", http.StatusBadRequest)
	}

	id := s.FileMetaDataStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s FileMetaDataResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.FileMetaDataStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s FileMetaDataResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(Model.FileMetaData)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid_Instance_Given"), "Invalid Instance Given", http.StatusBadRequest)
	}
	err := s.FileMetaDataStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
