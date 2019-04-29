package driver

import (
	"crypto/md5"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Storage in-memory map
type Storage struct {
	Store map[string]*BuildingData
	Lock  *sync.Mutex
}

// BuildingData data row in the storage
type BuildingData struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Floors   []string `json:"floors"`
	Created  string   `json:"created,omitempty"`
	Modified string   `json:"modified,omitempty"`
}

// NewStorage new storage object
func NewStorage() *Storage {
	return &Storage{
		Store: make(map[string]*BuildingData),
		Lock:  new(sync.Mutex),
	}
}

// Add new row
func (q *Storage) Add(name, address string, floors []string) (string, error) {

	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	// check if already in store
	storekey := q.StoreHash(name)
	row, oks := q.Store[storekey]
	if oks {
		//update
		row.Address = address
		row.Floors = floors
		row.Modified = time.Now().Format(time.RFC3339)
	} else {
		//new
		row = &BuildingData{
			ID:      storekey,
			Name:    name,
			Address: address,
			Floors:  floors,
			Created: time.Now().Format(time.RFC3339),
		}
	}
	//save it back
	q.Store[storekey] = row
	//give it back ;-)
	return storekey, nil
}

// Update an old record
func (q *Storage) Update(storekey, name, address string, floors []string) error {

	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	// check if id is a match
	if storekey != q.StoreHash(name) {
		//give it back ;-)
		return errors.New("record id mismatch ")
	}

	if row, oks := q.Store[storekey]; oks {
		//update
		row.Address = address
		row.Floors = floors
		row.Modified = time.Now().Format(time.RFC3339)
		//save it back
		q.Store[storekey] = row
		return nil
	}

	//give it back ;-)
	return errors.New("record not found")

}

// Delete an old record
func (q *Storage) Delete(storekey string) error {

	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if _, oks := q.Store[storekey]; oks {
		//delete
		delete(q.Store, storekey)
		return nil
	}

	//give it back ;-)
	return errors.New("record not found")

}

// GetOne 1 record
func (q *Storage) GetOne(storekey string) (*BuildingData, error) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	data, oks := q.Store[storekey]
	if !oks {
		return nil, errors.New("record not found")
	}

	//give it back ;-)
	return data, nil

}

// GetAll all the records
func (q *Storage) GetAll() ([]*BuildingData, error) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	var all []*BuildingData
	for _, row := range q.Store {
		all = append(all, row)
	}

	//empty
	if len(all) <= 0 {
		return all, errors.New("record not found")
	}
	//give it back ;-)
	return all, nil

}

// Len check total len
func (q *Storage) Len() int {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	//give it back ;-)
	return len(q.Store)
}

// StoreHash check if the name is in the storage
func (q *Storage) StoreHash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
