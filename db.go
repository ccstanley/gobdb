// Package gobdb is a put:1/get:1/list:n/delete:1 key:struct{} database using go's encoding/gob.
package gobdb

import "errors"

// Gobdb is a structure holding the db's data
type Gobdb struct {
	store  map[string]interface{}
	closed bool
}

// ErrClosed : The DB is already closed. No further action is allowed.
var ErrClosed = errors.New("DB already closed")

// OpenFile a gobdb file for read/write.
func OpenFile(filename string) (db *Gobdb, err error) {
	return nil, errors.New("NOT YET IMPLEMENTED")
}

// Open an in-memory gobdb database. This is not persistent in this case
func Open() (db *Gobdb, err error) {
	db = &Gobdb{
		store:  make(map[string]interface{}),
		closed: false,
	}
	return db, nil
}

// Close and flush the changes to the Gobdb file.
func (db *Gobdb) Close() error {
	db.closed = true
	return nil
}

// Put data into database with associated key
func (db *Gobdb) Put(key string, val interface{}) error {
	if db.closed {
		return ErrClosed
	}
	db.store[key] = val
	return nil
}

// Get an entry with a key
func (db *Gobdb) Get(key string) (interface{}, error) {
	if db.closed {
		return nil, ErrClosed
	}
	return db.store[key], nil
}

// List all data within the database. Note this is a resource heavy operation.
func (db *Gobdb) List() (ret map[string]interface{}, err error) {
	if db.closed {
		return nil, ErrClosed
	}
	// Ensure a new map instead of using the internal map.
	ret = make(map[string]interface{})

	for k, v := range db.store {
		ret[k] = v
	}
	return ret, nil
}

// Delete the value using a key.
func (db *Gobdb) Delete(key string) error {
	if db.closed {
		return ErrClosed
	}
	delete(db.store, key)
	return nil
}
