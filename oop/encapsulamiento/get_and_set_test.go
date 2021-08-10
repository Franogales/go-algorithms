package encapsulamiento_test

import (
	"testing"

	"github.com/Franogales/go-algorithms/oop/encapsulamiento"
)

/* probamos la instanciacion */
func TestNotName(t *testing.T) {
	auto := encapsulamiento.NewAuto()
	name := auto.GetName()

	if name != "" {
		t.Error("no se encontro name")
	}

}

func TestSettedName(t *testing.T) {
	auto := encapsulamiento.NewAuto()
	auto.SetName("malibu")
	name := auto.GetName()
	if name != "malibu" {
		t.Error("no se seteo el nombre del auto")
	}
}
