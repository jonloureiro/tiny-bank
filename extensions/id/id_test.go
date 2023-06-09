package id_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/extensions/id"
)

func TestNewID(t *testing.T) {
	id1 := id.New()
	id2 := id.New()
	if id1 == id2 {
		t.Errorf("expected id1 to be different from id2")
	}
}
