package gobdb

import (
	"io/ioutil"
	"os"
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

func TestGobdbPutGetDeleteGet(t *testing.T) {
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

	val, err := db.Get("aaa")
	if err != nil {
		t.Errorf("Error Get: %s", err)
		return
	}
	if val != "111" {
		t.Errorf("Got unexpected value: %+v", val)
		return
	}

	err = db.Delete("aaa")
	if err != nil {
		t.Errorf("Error deleting key from database: %s", err)
		return
	}

	val2, err := db.Get("aaa")
	if err != nil {
		t.Errorf("Error Get2: %s", err)
		return
	}
	if val2 != nil {
		t.Errorf("Got unexpected value: %+v, should be nil", val2)
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

func TestFileDbActions(t *testing.T) {
	f, err := ioutil.TempFile("", "gobdb")
	if err != nil {
		t.Errorf("Cannot create Tmp file for test: %s", err)
		return
	}
	fname := f.Name()
	defer os.Remove(fname)

	f.Close()

	db, err := OpenFile(fname)
	if err != nil {
		t.Errorf("Cannot open Tmp file for test: %s", err)
		return
	}

	db.Put("aaa", "111")
	db.Put("bb", 345)
	db.Close()

	// Open again and list the content
	db, _ = OpenFile(fname)
	testAssertDbContent(t, db, map[string]interface{}{
		"aaa": "111",
		"bb":  345,
	})

	db.Put("ccc", false)
	db.Close()

	db, _ = OpenFile(fname)
	testAssertDbContent(t, db, map[string]interface{}{
		"aaa": "111",
		"bb":  345,
		"ccc": false,
	})

	db.Delete("bb")
	db.Close()

	db, _ = OpenFile(fname)
	testAssertDbContent(t, db, map[string]interface{}{
		"aaa": "111",
		"ccc": false,
	})
	db.Close()
}

func testAssertDbContent(t *testing.T, db *Gobdb, expected map[string]interface{}) {
	res, err := db.List()
	if err != nil {
		t.Errorf("Error listing the database: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("The database is not returning correct data: %+v", res)
		return
	}
}
