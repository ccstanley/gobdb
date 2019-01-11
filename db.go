// Package gobdb is a put:1/get:1/list:n/delete:1 key:struct{} database using go's encoding/gob.
package gobdb

import "errors"

// Gobdb is a structure holding the db's data
type Gobdb struct {
}

// Open a gobdb file for read/write.
func Open(filename string) (db *Gobdb, err error) {
	return nil, errors.New("NOT YET IMPLEMENTED")
}

// Close and flush the changes to the Gobdb file.
func (db *Gobdb) Close() error {
	return errors.New("NOT IMPLEMENTED")
}

// Put data into database with associated key
func (db *Gobdb) Put(key string, val interface{}) error {
	return errors.New("NOT IMPLEMENTED")
}

// Get an entry with a key
func (db *Gobdb) Get(key string) (interface{}, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

// List all data within the database. Note this is a resource heavy operation.
func (db *Gobdb) List() (map[string]interface{}, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

// Delete the value using a key.
func (db *Gobdb) Delete(key string) error {
	return errors.New("NOT IMPLEMENTED")
}
