package main

import (
	"fmt"

	"github.com/Franogales/go-algorithms/oop/abstraccion"
	"github.com/Franogales/go-algorithms/oop/encapsulamiento"
)

func main() {
	exampleOfGetSet()
	exampleOfAbstraction()
	exampleOfEncapsulamiento()
}

func exampleOfAbstraction() {
	curso := abstraccion.Course{
		Name: "Example of abstraction",
		// price esta en minuscula dentro de la estructura y no podemos usarla fuera del packete 'abstraccion'
		IsFree:  true,
		UserIDs: []uint{1, 2, 3},
		Classes: map[uint]string{
			1: "CurseName1",
			2: "CurseName2",
			3: "CurseName3",
		},
	}
	fmt.Printf("%+v\n", curso)
}

/* esta funcion hace uso de una instanca de persona */
func exampleOfEncapsulamiento() {
	persona := encapsulamiento.NewPersona("Paco", "Nogales", 20)
	persona.Peso = 8
	fmt.Printf("%+v\n", persona)
}

/* de crean funciones especificas para manejar fuera del paquete los atributos de auto */
func exampleOfGetSet() {
	auto := encapsulamiento.NewAuto()
	auto.SetName("mersedes")
	fmt.Println(auto.GetName())
}
