package models

import (
	"context"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingDeleteParams delete parameter
type BuildingDeleteParams struct {
	ID string `json:"ID,required"`
}

// NewBuildingDelete data remover
func NewBuildingDelete(pid string) *BuildingDeleteParams {
	return &BuildingDeleteParams{ID: pid}
}

// Remove delete a row from the store base on id
func (p *BuildingDeleteParams) Remove(ctx context.Context, store *drivers.Storage) error {
	if _, oks := store.Exists(p.ID); !oks {
		return ErrRecordNotFound
	}
	return store.Delete(p.ID)
}
