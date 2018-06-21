package saveVars

import (
	"testing"
)

func TestVarsInt(t *testing.T) {
	InitVars()
	SetInt(100500, "test")
	val := GetInt("test")
	if val != 100500 {
		t.Fatal("Values not match")
	}
	t.Log("Value of `test` var: ", val)
}
