package models

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/bayugyug/rest-building/driver"
)

// BuildingUpdateParams update params
type BuildingUpdateParams struct {
	ID string `json:"id,required"`
	BuildingCreateParams
}

// NewBuildingUpdate new creator
func NewBuildingUpdate() *BuildingUpdateParams {
	return &BuildingUpdateParams{}

}

// Bind filter params
func (params *BuildingUpdateParams) Bind(r *http.Request) error {
	//sanity check
	if params == nil {
		return errors.New("missing required parameter")
	}
	//check id
	params.ID = strings.TrimSpace(params.ID)
	params.Name = strings.TrimSpace(params.Name)
	params.Address = strings.TrimSpace(params.Address)
	if params.Name == "" || params.Address == "" || len(params.Floors) == 0 || params.ID == "" {
		return errors.New("missing required parameter")
	}

	// just a post-process after a decode..
	return nil
}

// Update a row from the store
func (params *BuildingUpdateParams) Update(ctx context.Context, store *driver.Storage) error {
	return store.Update(params.ID, params.Name, params.Address, params.Floors)
}
