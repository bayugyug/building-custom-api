package models

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingUpdateParams update parameter
type BuildingUpdateParams struct {
	ID string `json:"id,required"`
	BuildingCreateParams
}

// NewBuildingUpdate new creator
func NewBuildingUpdate() *BuildingUpdateParams {
	return &BuildingUpdateParams{}
}

// Bind filter parameter
func (p *BuildingUpdateParams) Bind(r *http.Request) error {
	//sanity check
	if p == nil {
		return ErrMissingRequiredParameters
	}
	//check
	p.ID = strings.TrimSpace(p.ID)
	p.Name = strings.TrimSpace(p.Name)
	p.Address = strings.TrimSpace(p.Address)
	if !p.SanityCheck() {
		return ErrMissingRequiredParameters
	}
	// just a post-process after a decode..
	return nil
}

// SanityCheck filter required parameter
func (p *BuildingUpdateParams) SanityCheck() bool {
	if p.Name == "" || p.Address == "" ||
		len(p.Floors) == 0 || p.ID == "" {
		return false
	}
	return true
}

// Update a row from the store
func (p *BuildingUpdateParams) Update(ctx context.Context, store *drivers.Storage) error {
	//check
	if !p.SanityCheck() {
		return ErrMissingRequiredParameters
	}
	//db check
	row, oks := store.Exists(p.ID)
	if !oks {
		return ErrRecordNotFound
	}

	//check the hashkey
	record := NewBuildingData()
	pid := record.HashKey(p.Name)
	if pid != p.ID {
		log.Println(pid, "==================", p.ID, "==", p.Name)
		return ErrRecordMismatch
	}
	//convert db data
	vrow, ok := row.(*BuildingData)
	if !ok {
		return ErrRecordNotFound
	}
	//set row
	record = vrow
	record.Name = p.Name
	record.Address = p.Address
	record.Floors = p.Floors
	record.Modified = time.Now().Format(time.RFC3339)
	_, err := store.Set(record.ID, record)
	return err
}
