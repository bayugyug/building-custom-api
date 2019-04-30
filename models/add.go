package models

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/driver"
)

var (
	// ErrMissingRequiredParameters reqd params missing
	ErrMissingRequiredParameters = errors.New("missing required parameter")
	// ErrorNotFound record not found
	ErrorNotFound = errors.New("record not found")
	// ErrRecordsNotFound list is empty
	ErrRecordsNotFound = errors.New("record(s) not found")
	// ErrRecordNotFound data not exiss
	ErrRecordNotFound = errors.New("record not found")
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
		return ErrMissingRequiredParameters
	}
	params.Name = strings.TrimSpace(params.Name)
	params.Address = strings.TrimSpace(params.Address)

	//check
	if !params.SanityCheck() {
		return ErrMissingRequiredParameters
	}

	// just a post-process after a decode..
	return nil
}

// SanityCheck filter required params
func (params *BuildingCreateParams) SanityCheck() bool {
	if params.Name == "" || params.Address == "" ||
		len(params.Floors) == 0 {
		return false
	}
	return true
}

// Create add a row from the store
func (params *BuildingCreateParams) Create(ctx context.Context, store *driver.Storage) (string, error) {
	//check
	if !params.SanityCheck() {
		return "", ErrMissingRequiredParameters
	}

	record := driver.NewBuildingData()
	pid := record.StoreKey(params.Name)

	if oks := store.Exists(pid); oks {
		record.Modified = time.Now().Format(time.RFC3339)
	} else {
		record.Created = time.Now().Format(time.RFC3339)
	}
	//set row
	record.ID = pid
	record.Name = params.Name
	record.Address = params.Address
	record.Floors = params.Floors
	return store.Set(pid, record)
}
