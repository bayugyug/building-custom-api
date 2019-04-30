package driver

import (
	"errors"
	"sync"
)

var (
	// ErrEmptyParameter some required params empty
	ErrEmptyParameter = errors.New("record name/id empty")
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
func (q *Storage) Set(storekey string, data interface{}) (string, error) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	// check
	if len(storekey) == 0 {
		return "", ErrEmptyParameter
	}
	// check if already in store
	q.Store[storekey] = data
	return storekey, nil
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
	return ErrRecordNotFound
}

// GetOne 1 record
func (q *Storage) GetOne(storekey string) (interface{}, error) {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()

	data, oks := q.Store[storekey]
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
func (q *Storage) Exists(storekey string) bool {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	_, oks := q.Store[storekey]
	//give it back ;-)
	return oks
}

// Len check total len
func (q *Storage) Len() int {
	// ensure
	q.Lock.Lock()
	defer q.Lock.Unlock()
	//give it back ;-)
	return len(q.Store)
}
