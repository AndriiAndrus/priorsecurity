package datastore

import (
	"testing"
)

const file = "../test_file.gob"

type User struct {
	Name, Pass string
}

func TestGoData(t *testing.T) {
	var datato = &User{"Donald", "DuckPass"}
	var datafrom = new(User)

	err := Save(file, datato)
	if err != nil {
		t.Error(err)
	}
	err = Load(file, datafrom)
	if err != nil {
		t.Error(err)
	}
	t.Log(datafrom)
}
