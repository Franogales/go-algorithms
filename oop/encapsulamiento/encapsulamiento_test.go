package encapsulamiento_test

import (
	"testing"

	"github.com/Franogales/go-algorithms/oop/encapsulamiento"
)

func TestConstructor(t *testing.T) {
	persona := encapsulamiento.NewPersona("", "", 0)
	if persona == nil {
		t.Errorf("pesona is empty")
	}
}
