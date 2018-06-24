package saveVars

import (
	"math/rand"
	"mobiSec/crypto"
	"testing"
)

func TestVarsStr(t *testing.T) {
	var ds = GetDatastore(crypto.NewEncryptionKey(), "/workspaces/AndriiAndrus/src/mobiSec/vars.gob")

	ds.SetString("Test string", "Test string nadibf jfgb 78464")
	ds.SetInt("Test int", rand.Int())
	ds.SetFloat("Test float", rand.Float64())

	ds.save()

	str, err := ds.GetString("Test string")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Value of `Test string` var: ", str)

	in, err := ds.GetInt("Test int")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Value of `Test int` var: ", in)

	fl, err := ds.GetFloat("Test float")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Value of `Test float`: ", fl)
}
