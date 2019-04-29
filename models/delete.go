package models

import (
	"context"

	"github.com/bayugyug/rest-building/driver"
)

// BuildingDeleteParams delete params
type BuildingDeleteParams struct {
	ID string `json:"ID,required"`
}

// NewBuildingDelete data remover
func NewBuildingDelete(pid string) *BuildingDeleteParams {
	return &BuildingDeleteParams{ID: pid}

}

// Remove delete a row from the store base on id
func (params *BuildingDeleteParams) Remove(ctx context.Context, store *driver.Storage) error {
	return store.Delete(params.ID)
}
