package drivers

import (
	"errors"
	"sync"
)

var (
	// ErrRecordNotFound not exists
	ErrRecordNotFound = errors.New("record not found")
)

// Storage in-memory map
type Storage struct {
	Store map[string]interface{}
	Lock  *sync.Mutex
}

// NewStorage new storage object
func NewStorage() *Storage {
	return &Storage{
		Store: make(map[string]interface{}),
		Lock:  new(sync.Mutex),
	}
}

// Set new row
func (q *Storage) Set(key string, data interface{}) string {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	q.Store[key] = data
	return key
}

// Delete an old record
func (q *Storage) Delete(key string) error {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if _, oks := q.Store[key]; oks {
		//delete
		delete(q.Store, key)
		return nil
	}
	//give it back ;-)
	return ErrRecordNotFound
}

// GetOne 1 record
func (q *Storage) GetOne(key string) (interface{}, error) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	data, oks := q.Store[key]
	if !oks {
		return nil, ErrRecordNotFound
	}
	//give it back ;-)
	return data, nil
}

// GetAll all the records
func (q *Storage) GetAll() ([]interface{}, error) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	var all []interface{}
	for _, row := range q.Store {
		all = append(all, row)
	}
	//give it back ;-)
	return all, nil
}

// Exists check the record
func (q *Storage) Exists(key string) (interface{}, bool) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	row, oks := q.Store[key]
	//give it back ;-)
	return row, oks
}

// Len check total len
func (q *Storage) Len() int {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	//give it back ;-)
	return len(q.Store)
}
