package models

import (
	"context"

	"github.com/bayugyug/building-custom-api/driver"
)

// BuildingGetOneParams delete params
type BuildingGetOneParams struct {
	ID string `json:"ID"`
}

// NewBuildingGetOne data remover
func NewBuildingGetOne(id string) *BuildingGetOneParams {
	return &BuildingGetOneParams{ID: id}
}

// Get query from the store base on id
func (params *BuildingGetOneParams) Get(ctx context.Context, store *driver.Storage) (*driver.BuildingData, error) {
	data, err := store.GetOne(params.ID)
	if err != nil {
		return nil, err
	}

	var rec *driver.BuildingData
	var valid bool

	if rec, valid = data.(*driver.BuildingData); valid {
		return rec, nil
	}
	//not found
	return nil, ErrRecordsNotFound
}

// GetAll query from the store base on id
func (params *BuildingGetOneParams) GetAll(ctx context.Context, store *driver.Storage) ([]*driver.BuildingData, error) {
	data, err := store.GetAll()
	if err != nil {
		return nil, err
	}

	var all []*driver.BuildingData

	//empty
	if len(data) <= 0 {
		return all, ErrRecordsNotFound
	}

	for _, vv := range data {
		if row, valid := vv.(*driver.BuildingData); valid {
			all = append(all, row)
		}
	}
	return all, nil
}
