package models

import (
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
)

// BuildingUpdateParams update parameter
type BuildingUpdateParams struct {
	ID *string `json:"id"`
	BuildingCreateParams
}

// NewBuildingUpdate new instance
func NewBuildingUpdate() *BuildingUpdateParams {
	return &BuildingUpdateParams{}
}

// Bind filter parameter
func (p *BuildingUpdateParams) Bind(r *http.Request) error {
	//sanity check
	if p == nil {
		return ErrMissingRequiredParameters
	}
	//fmt
	p.Address = strings.TrimSpace(p.Address)
	//chk
	return p.SanityCheck()
}

// SanityCheck filter required parameter
func (p *BuildingUpdateParams) SanityCheck() error {
	if p.ID == nil || p.Name == nil ||
		*p.ID == "" || *p.Name == "" {
		return ErrMissingRequiredParameters
	}
	return nil
}

// Update a row from the store
func (p *BuildingUpdateParams) Update(store *drivers.Storage) error {
	//should not happen :-)
	if err := p.SanityCheck(); err != nil {
		return err
	}
	//db check
	row, oks := store.Exists(*p.ID)
	if !oks {
		return ErrRecordNotFound
	}
	//check the hashkey
	record := NewBuildingData()
	pid := record.HashKey(*p.Name)
	if pid != *p.ID {
		return ErrRecordMismatch
	}
	//convert db data
	vrow, ok := row.(*BuildingData)
	if !ok {
		return ErrDBTransaction
	}
	//set old row with new value
	record = vrow
	record.Address = p.Address
	record.Floors = p.Floors
	record.Modified = time.Now().Format(time.RFC3339)
	if gid := store.Set(pid, record); gid == "" {
		return ErrDBTransaction
	}
	return nil
}
