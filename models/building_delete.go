package models

import (
	"context"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingDeleteParams delete parameter
type BuildingDeleteParams struct {
	ID string `json:"id"`
}

// NewBuildingDelete new instance
func NewBuildingDelete(pid string) *BuildingDeleteParams {
	return &BuildingDeleteParams{ID: pid}
}

// Delete remove a row from the store base on id
func (p *BuildingDeleteParams) Delete(ctx context.Context, store *drivers.Storage) error {
	if _, oks := store.Exists(p.ID); !oks {
		return ErrRecordNotFound
	}
	return store.Unset(p.ID)
}
