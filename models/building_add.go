package models

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
)

var (
	// ErrMissingRequiredParameters reqd parameter missing
	ErrMissingRequiredParameters = errors.New("missing required parameter")
	// ErrorNotFound record not found
	ErrorNotFound = errors.New("record not found")
	// ErrRecordsNotFound list is empty
	ErrRecordsNotFound = errors.New("record(s) not found")
	// ErrRecordNotFound data not exiss
	ErrRecordNotFound = errors.New("record not found")
	// ErrRecordMismatch generated hashkey by name is a mismatch
	ErrRecordMismatch = errors.New("record id/name mismatch")
	// ErrRecordExists data already exiss
	ErrRecordExists = errors.New("record exists")
)

// BuildingCreateParams create parameter
type BuildingCreateParams struct {
	Name    string   `json:"name,required"`
	Address string   `json:"address,required"`
	Floors  []string `json:"floors,required"`
}

// NewBuildingCreate new creator
func NewBuildingCreate() *BuildingCreateParams {
	return &BuildingCreateParams{}

}

// Bind filter parameter
func (p *BuildingCreateParams) Bind(r *http.Request) error {
	//sanity check
	if p == nil {
		return ErrMissingRequiredParameters
	}
	p.Name = strings.TrimSpace(p.Name)
	p.Address = strings.TrimSpace(p.Address)

	//check
	if !p.SanityCheck() {
		return ErrMissingRequiredParameters
	}

	// just a post-process after a decode..
	return nil
}

// SanityCheck filter required parameter
func (p *BuildingCreateParams) SanityCheck() bool {
	if p.Name == "" || p.Address == "" ||
		len(p.Floors) == 0 {
		return false
	}
	return true
}

// Create add a row from the store
func (p *BuildingCreateParams) Create(ctx context.Context, store *drivers.Storage) (string, error) {
	//should not happen
	if !p.SanityCheck() {
		return "", ErrMissingRequiredParameters
	}
	record := NewBuildingData()
	pid := record.HashKey(p.Name)
	if _, oks := store.Exists(pid); oks {
		return "", ErrRecordExists
	}
	//set row
	record.ID = pid
	record.Created = time.Now().Format(time.RFC3339)
	record.Name = p.Name
	record.Address = p.Address
	record.Floors = p.Floors
	return store.Set(pid, record)
}
