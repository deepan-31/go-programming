package ap2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("Hello is not email")
	}
	_, err = IsEmail("derek@aol.com")
	if err == nil {
		t.Error("derek@aol.com is an email")
	}
	_, err = IsEmail("derekaol.com")
	if err != nil {
		t.Error("derekaol.com is not an email")
	}
}
