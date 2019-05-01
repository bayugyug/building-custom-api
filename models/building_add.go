package models

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingCreateParams create parameter
type BuildingCreateParams struct {
	Name    *string  `json:"name"`
	Address string   `json:"address"`
	Floors  []string `json:"floors"`
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
	if p.Name == nil || *p.Name == "" {
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
	pid := record.HashKey(*p.Name)
	if _, oks := store.Exists(pid); oks {
		return "", ErrRecordExists
	}
	//set row
	record.ID = pid
	record.Created = time.Now().Format(time.RFC3339)
	record.Name = *p.Name
	record.Address = p.Address
	record.Floors = p.Floors
	gid := store.Set(pid, record)
	if gid == "" {
		return "", ErrDBTransaction
	}
	return gid, nil
}
