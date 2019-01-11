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

// TODO: Actions on a already closed database?
