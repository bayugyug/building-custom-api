package models

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/bayugyug/building-custom-api/driver"
)

// BuildingCreateParams create params
type BuildingCreateParams struct {
	Name    string   `json:"name,required"`
	Address string   `json:"address,required"`
	Floors  []string `json:"floors,required"`
}

// NewBuildingCreate new creator
func NewBuildingCreate() *BuildingCreateParams {
	return &BuildingCreateParams{}

}

// Bind filter params
func (params *BuildingCreateParams) Bind(r *http.Request) error {
	//sanity check
	if params == nil {
		return errors.New("missing required parameter")
	}
	//check
	params.Name = strings.TrimSpace(params.Name)
	params.Address = strings.TrimSpace(params.Address)
	if params.Name == "" || params.Address == "" ||
		len(params.Floors) == 0 {
		return errors.New("missing required parameter")
	}
	// just a post-process after a decode..
	return nil
}

// Create add a row from the store
func (params *BuildingCreateParams) Create(ctx context.Context, store *driver.Storage) (string, error) {
	return store.Add(params.Name, params.Address, params.Floors)
}
