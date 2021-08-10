package abstraccion_test

import (
	"testing"

	"github.com/Franogales/go-algorithms/oop/abstraccion"
)

func TestAbstraccion(t *testing.T) {
	curse := abstraccion.Course{
		Name: "olis",
	}
	if curse.Name != "olis" {
		t.Errorf("Name not work")
	}
}
