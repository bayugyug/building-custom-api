package models

import (
	"crypto/md5"
	"errors"
	"fmt"
)

var (
	// ErrMissingRequiredParameters reqd parameter missing
	ErrMissingRequiredParameters = errors.New("missing required parameter")
	// ErrRecordsNotFound list is empty
	ErrRecordsNotFound = errors.New("record(s) not found")
	// ErrRecordNotFound data not exiss
	ErrRecordNotFound = errors.New("record not found")
	// ErrRecordMismatch generated hashkey by name is a mismatch
	ErrRecordMismatch = errors.New("record id/name mismatch")
	// ErrRecordExists data already exiss
	ErrRecordExists = errors.New("record exists")
	// ErrDBTransaction internal storage error
	ErrDBTransaction = errors.New("db storage failed")
)

// BuildingData data row in the storage
type BuildingData struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address,omitempty"`
	Floors   []string `json:"floors,omitempty"`
	Created  string   `json:"created,omitempty"`
	Modified string   `json:"modified,omitempty"`
}

// NewBuildingData new storage object
func NewBuildingData() *BuildingData {
	return &BuildingData{}
}

// HashKey convert to md5 hash
func (q BuildingData) HashKey(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
