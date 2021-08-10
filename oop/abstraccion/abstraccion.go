package abstraccion

import "fmt"

/*
este programa solo demostrara la abstraccion a travez de las estructuras.
En la abstraccion exportamos los atributos que queremos exportar en el packete
En este caso vamos a evitar que los demas paquetes puedan utilizar el atributo price de la structura Course
*/
// Course is the struct that contains all the
type Course struct {
	Name    string
	price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// en la funccion print imprimimos todos los elementos de Course seteados arbitrareamente
// solo para demostrar que si podemos utilizar todos los atributos de Course dentro del paquete pero no por fuera
func Define() {
	curso := Course{
		Name:    "Go OOP",
		price:   10.0,
		IsFree:  true,
		UserIDs: []uint{1, 2, 3},
		Classes: map[uint]string{
			1: "CurseName1",
			2: "CurseName2",
			3: "CurseName3",
		},
	}
	fmt.Printf(curso.Name)
}
