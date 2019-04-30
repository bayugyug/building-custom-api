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
	return data, nil
}

// GetAll query from the store base on id
func (params *BuildingGetOneParams) GetAll(ctx context.Context, store *driver.Storage) ([]*driver.BuildingData, error) {
	if data, err := store.GetAll(); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
