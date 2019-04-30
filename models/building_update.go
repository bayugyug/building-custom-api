package models

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/driver"
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
		return ErrMissingRequiredParameters
	}
	//check
	params.ID = strings.TrimSpace(params.ID)
	params.Name = strings.TrimSpace(params.Name)
	params.Address = strings.TrimSpace(params.Address)
	if !params.SanityCheck() {
		return ErrMissingRequiredParameters
	}
	// just a post-process after a decode..
	return nil
}

// SanityCheck filter required params
func (params *BuildingUpdateParams) SanityCheck() bool {
	if params.Name == "" || params.Address == "" ||
		len(params.Floors) == 0 || params.ID == "" {
		return false
	}
	return true
}

// Update a row from the store
func (params *BuildingUpdateParams) Update(ctx context.Context, store *driver.Storage) error {
	//check
	if !params.SanityCheck() {
		return ErrMissingRequiredParameters
	}
	if oks := store.Exists(params.ID); !oks {
		return ErrRecordNotFound
	}
	//set row
	record := &driver.BuildingData{
		ID:       params.ID,
		Name:     params.Name,
		Address:  params.Address,
		Floors:   params.Floors,
		Modified: time.Now().Format(time.RFC3339),
	}
	_, err := store.Set(record.ID, record)
	return err
}
