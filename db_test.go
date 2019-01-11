package gobdb

import (
	"reflect"
	"testing"
)

func TestGobdbOpenClose(t *testing.T) {
	db, err := Open()
	if err != nil {
		t.Errorf("Cannot open the Memory database: %s", err)
		return
	}
	err = db.Close()
	if err != nil {
		t.Errorf("Cannot close database: %s", err)
	}
}
func TestGobdbPutAndList(t *testing.T) {
	db, _ := Open()
	defer db.Close()

	testdata := map[string]interface{}{
		"aaa": "111",
		"bb":  345,
	}

	for k, v := range testdata {
		if err := db.Put(k, v); err != nil {
			t.Errorf("Error inserting database data: k=%s, v=%+v, err=%s", k, v, err)
			return
		}
	}

	res, err := db.List()
	if err != nil {
		t.Errorf("Error listing database: %s", err)
		return
	}

	if !reflect.DeepEqual(testdata, res) {
		t.Errorf("The database is not returning correct data.")
		return
	}
}

func TestActOnClosedDb(t *testing.T) {
	db, _ := Open()

	testdata := map[string]interface{}{
		"aaa": "111",
		"bb":  345,
	}

	for k, v := range testdata {
		db.Put(k, v)
	}

	db.Close()

	if err := db.Put("cc", 123); err != ErrClosed {
		t.Errorf("Put should return error")
	}
	if _, err := db.Get("aaa"); err != ErrClosed {
		t.Errorf("Get should return error")
	}
	if _, err := db.List(); err != ErrClosed {
		t.Errorf("List should return error")
	}
	if err := db.Delete("bb"); err != ErrClosed {
		t.Errorf("Delete should return error")
	}
}
