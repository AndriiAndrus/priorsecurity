package saveVars

import (
	"encoding/base64"
	"fmt"
	"mobiSec/crypto"
	"mobiSec/datastore"
	"os"
	"strconv"
)

type Backup struct {
	vars map[string][]byte
}

type Datastore struct {
	vars     map[string][]byte
	file     string
	password *[32]byte
}

// Create or load the instance of Datastore
func GetDatastore(key *[32]byte, path string) *Datastore {
	if _, err := os.Stat(path); err == nil {
		//var x *Backup // px is initialized to nil.
		x := new(Backup)
		x.vars = make(map[string][]byte)
		err := datastore.Load(path, x)
		if err != nil {
			fmt.Println("Datastore restore error: ", err)
		}

		return &Datastore{x.vars, path, key}
	}

	var m = make(map[string][]byte)

	return &Datastore{m, path, key}
}

// Save current datastore vars to the disk
func (d *Datastore) save() {
	back := new(Backup)
	back.vars = d.vars
	err := datastore.Save(d.file, back)
	if err != nil {
		panic(err)
	}
}

func (d *Datastore) SetFloat(name string, val float64) error {
	str := strconv.FormatFloat(val, 'f', -1, 64)
	err := d.set([]byte(str), name)

	return err
}

func (d *Datastore) GetFloat(name string) (float64, error) {
	bytes, err := d.get(name)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	float, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		return 0, err
	}

	return float, nil
}

func (d *Datastore) SetString(name string, val string) error {
	err := d.set([]byte(val), name)

	return err
}

func (d *Datastore) GetString(name string) (string, error) {
	bytes, err := d.get(name)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}

	return string(bytes), nil
}

func (d *Datastore) SetInt(name string, val int) error {
	converted := strconv.Itoa(val)

	err := d.set([]byte(converted), name)

	return err
}

func (d *Datastore) GetInt(name string) (int, error) {
	bytes, err := d.get(name)
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

func (d *Datastore) set(val []byte, name string) error {
	dataBase64 := base64.StdEncoding.EncodeToString(val)

	encrypted, err := crypto.Encrypt([]byte(dataBase64), d.password)
	if err != nil {
		fmt.Println("Vars encryption error: ", err)
		return err
	}

	d.vars[name] = encrypted

	return nil
}

func (d *Datastore) get(name string) ([]byte, error) {
	decrypted, err := crypto.Decrypt(d.vars[name], d.password)
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
