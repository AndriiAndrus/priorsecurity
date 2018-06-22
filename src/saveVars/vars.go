package saveVars

import (
	"encoding/base64"
	"fmt"
	"mobiSec/crypto"
	"mobiSec/datastore"
	"os"
	"strconv"
)

type SecuredData struct {
	values map[string][]byte
}

var file string
var datafrom *SecuredData
var password *[32]byte

func InitVars(key *[32]byte, pathToDatastore string) {
	if file == "" && password == nil {
		file = pathToDatastore
		password = key
	}

	if _, err := os.Stat(file); err == nil {
		// datastore file exist, lets try to load it
		_, e := load()
		if e != nil {
			datafrom = new(SecuredData)
			datafrom.values = make(map[string][]byte)

			err := save(datafrom)
			if err != nil {
				fmt.Println("Error init vars: ", err)
			}
		}
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

func SetString(val string, name string) error {
	err := set([]byte(val), name)

	return err
}

func GetString(name string) (string, error) {
	bytes, err := get(name)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	return string(bytes), nil
}

func SetInt(val int, name string) error {
	converted := strconv.Itoa(val)

	err := set([]byte(converted), name)

	return err
}

func GetInt(name string) (int, error) {
	bytes, err := get(name)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	number, err := strconv.Atoi(string(bytes))
	if err != nil {
		fmt.Println("Error converting val to number: ", err)
		return 0, err
	}

	return number, nil
}

func set(val []byte, name string) error {
	datafrom, err := load()
	if err != nil {
		fmt.Println("Error loading vars: ", err)
		return err
	}

	dataBase64 := base64.StdEncoding.EncodeToString(val)

	encrypted, err := crypto.Encrypt([]byte(dataBase64), password)
	if err != nil {
		fmt.Println("Vars encryption error: ", err)
		return err
	}

	datafrom.values[name] = encrypted

	err = save(datafrom)
	if err != nil {
		fmt.Println("Error saving vars: ", err)
		return err
	}

	return nil
}

func get(name string) ([]byte, error) {
	datafrom, err := load()
	if err != nil {
		fmt.Println("Error loading vars: ", err)
		return nil, err
	}

	decrypted, err := crypto.Decrypt(datafrom.values[name], password)
	if err != nil {
		fmt.Println("Var decryption error: ", err)
		return nil, err
	}

	bytes, err := base64.StdEncoding.DecodeString(string(decrypted))
	if err != nil {
		fmt.Println("Error converting val from base64: ", err)
		return nil, err
	}

	return bytes, nil
}
