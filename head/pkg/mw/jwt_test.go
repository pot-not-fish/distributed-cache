package mw

import (
	"testing"
)

func TestJwt(t *testing.T) {
	jwt, err := GenerateJWT()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if err := ParseToken(jwt); err != nil {
		t.Fatalf(err.Error())
	}
}
