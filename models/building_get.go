package models

import (
	"context"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingGetOneParams delete parameter
type BuildingGetOneParams struct {
	ID string `json:"id"`
}

// NewBuildingGetOne data remover
func NewBuildingGetOne(id string) *BuildingGetOneParams {
	return &BuildingGetOneParams{ID: id}
}

// Get query from the store base on id
func (p *BuildingGetOneParams) Get(ctx context.Context, store *drivers.Storage) (*BuildingData, error) {
	data, err := store.GetOne(p.ID)
	if err != nil {
		return nil, err
	}
	var rec *BuildingData
	var valid bool
	if rec, valid = data.(*BuildingData); valid {
		return rec, nil
	}
	//not found
	return nil, ErrRecordsNotFound
}

// GetAll query from the store base on id
func (p *BuildingGetOneParams) GetAll(ctx context.Context, store *drivers.Storage) ([]*BuildingData, error) {
	data, err := store.GetAll()
	if err != nil {
		return nil, err
	}
	var all []*BuildingData
	//empty
	if len(data) <= 0 {
		return all, ErrRecordsNotFound
	}
	for _, vv := range data {
		if row, valid := vv.(*BuildingData); valid {
			all = append(all, row)
		}
	}
	return all, nil
}
