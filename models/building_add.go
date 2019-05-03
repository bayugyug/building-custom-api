package models

import (
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
	return p.SanityCheck()
}

// SanityCheck filter required parameter
func (p *BuildingCreateParams) SanityCheck() error {
	if p.Name == nil || *p.Name == "" {
		return ErrMissingRequiredParameters
	}
	return nil
}

// Create add a row from the store
func (p *BuildingCreateParams) Create(store *drivers.Storage) (string, error) {
	//should not happen
	if err := p.SanityCheck(); err != nil {
		return "", err
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
