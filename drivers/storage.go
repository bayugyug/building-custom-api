package drivers

import (
	"errors"
	"sync"
)

var (
	// ErrRecordNotFound not found record
	ErrRecordNotFound = errors.New("record not found")
)

// StorageDriver in-memory map
type StorageDriver interface {
	Set(key string, data interface{}) string
	Unset(key string) error
	One(key string) (interface{}, error)
	All() ([]interface{}, error)
}

// Storage in-memory map
type Storage struct {
	store map[string]interface{}
	mtx   *sync.Mutex
}

// NewStorage new storage object
func NewStorage() *Storage {
	return &Storage{
		store: make(map[string]interface{}),
		mtx:   new(sync.Mutex),
	}
}

// Set new row
func (q *Storage) Set(key string, data interface{}) string {
	// ensure
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.store[key] = data
	return key
}

// Unset an old record
func (q *Storage) Unset(key string) error {
	// ensure
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if _, oks := q.store[key]; oks {
		//delete
		delete(q.store, key)
		return nil
	}
	//give it back ;-)
	return ErrRecordNotFound
}

// One get 1 record
func (q *Storage) One(key string) (interface{}, error) {
	// ensure
	q.mtx.Lock()
	defer q.mtx.Unlock()

	data, oks := q.store[key]
	if !oks {
		return nil, ErrRecordNotFound
	}
	//give it back ;-)
	return data, nil
}

// All get list of all the records
func (q *Storage) All() ([]interface{}, error) {
	// ensure
	q.mtx.Lock()
	defer q.mtx.Unlock()

	var all []interface{}
	for _, row := range q.store {
		all = append(all, row)
	}
	//give it back ;-)
	return all, nil
}

// Exists check the record
func (q *Storage) Exists(key string) (interface{}, bool) {
	// ensure
	q.mtx.Lock()
	defer q.mtx.Unlock()
	row, oks := q.store[key]
	//give it back ;-)
	return row, oks
}

// Count check total len
func (q *Storage) Count() int {
	// ensure
	q.mtx.Lock()
	defer q.mtx.Unlock()
	//give it back ;-)
	return len(q.store)
}
