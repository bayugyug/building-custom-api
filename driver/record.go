package driver

import (
	"crypto/md5"
	"fmt"
)

// BuildingData data row in the storage
type BuildingData struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Floors   []string `json:"floors"`
	Created  string   `json:"created,omitempty"`
	Modified string   `json:"modified,omitempty"`
}

// NewBuildingData new storage object
func NewBuildingData() *BuildingData {
	return &BuildingData{}
}

// StoreKey convert to md5 hash
func (q *BuildingData) StoreKey(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
