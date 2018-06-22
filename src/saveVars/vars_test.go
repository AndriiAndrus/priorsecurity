package saveVars

import (
	"math/rand"
	"mobiSec/crypto"
	"testing"
)

func TestVarsStr(t *testing.T) {
	InitVars(crypto.NewEncryptionKey(), "/workspaces/AndriiAndrus/src/mobiSec/vars.gob")

	value := "Spicy Chongqing hogo is amazing!"

	err := SetString(value, "testStr")
	if err != nil {
		t.Fatal(err)
	}
	val, err := GetString("testStr")
	if err != nil {
		t.Fatal(err)
	}
	if val != value {
		t.Fatal("Values not match")
	}
	t.Log("Value of `testStr` var: ", val)
}

func TestVarsInt(t *testing.T) {
	//InitVars(crypto.NewEncryptionKey(), "/workspaces/AndriiAndrus/src/mobiSec/vars.gob")

	value := rand.Int()

	err := SetInt(value, "test")
	if err != nil {
		t.Fatal(err)
	}
	val, err := GetInt("test")
	if err != nil {
		t.Fatal(err)
	}
	if val != value {
		t.Fatal("Values not match")
	}
	t.Log("Value of `test` var: ", val)
}

func TestLoadVars(t *testing.T) {
	val, err := GetString("testStr")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Loaded string: ", val)

	val2, err := GetInt("test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Loaded int: ", val2)
}
