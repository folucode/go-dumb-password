package dumbpwd

import (
	"testing"
)

func TestDumbpwd(t *testing.T) {
	pwd := "hello"

	msg, err := CheckPassword(pwd)

	if err {
		t.Log(msg)
	}
}
