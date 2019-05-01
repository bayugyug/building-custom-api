package models

import (
	"context"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingGetParams get parameter
type BuildingGetParams struct {
	ID string `json:"id"`
}

// NewBuildingGetOne new instance with parameter
func NewBuildingGetOne(id string) *BuildingGetParams {
	return &BuildingGetParams{ID: id}
}

// Get query from the store base on id
func (p *BuildingGetParams) Get(ctx context.Context, store *drivers.Storage) (*BuildingData, error) {
	data, err := store.One(p.ID)
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
func (p *BuildingGetParams) GetAll(ctx context.Context, store *drivers.Storage) ([]*BuildingData, error) {
	data, err := store.All()
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
