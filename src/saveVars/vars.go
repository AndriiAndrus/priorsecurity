package saveVars

import (
	"fmt"
	"mobiSec/datastore"
)

type SecuredData struct {
	ints map[string]int
}

const file = "/workspaces/AndriiAndrus/src/mobiSec/vars.gob"

var datafrom *SecuredData

func InitVars() {
	datafrom = new(SecuredData)
	datafrom.ints = make(map[string]int)
	err := save(datafrom)
	if err != nil {
		fmt.Println("Error init vars: ", err)
	}
}

func save(datato *SecuredData) error {
	err := datastore.Save(file, datato)
	if err != nil {
		return err
	}
	return nil
}

func load() (*SecuredData, error) {
	if datafrom != nil {
		return datafrom, nil
	}
	datafrom = new(SecuredData)
	err := datastore.Load(file, datafrom)
	if err != nil {
		return nil, err
	}
	return datafrom, nil
}

func SetInt(val int, name string) {
	datafrom, err := load()
	if err != nil {
		fmt.Println("Error loading vars: ", err)
	}
	datafrom.ints[name] = val
	err = save(datafrom)
	if err != nil {
		fmt.Println("Error saving vars: ", err)
	}
}

func GetInt(name string) int {
	datafrom, err := load()
	if err != nil {
		fmt.Println("Error loading vars: ", err)
	}
	return datafrom.ints[name]
}
