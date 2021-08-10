/*  Lo importante en el encapsulamiento es que solamente
permitamos el acceso a lo que realmente es exportable
*/

package encapsulamiento

/* la estructura persona solamente puede ser accedida dentro del paquete encapsulamiento */
type persona struct {
	Name     string
	LastName string
	Edad     int
	Peso     int
}

/* con la palabra New en la funcion es como definimos el constructor en go (no existen constructores) */
func NewPersona(name, lastName string, edad int) *persona {
	if edad == 0 {
		edad = 1
	}
	return &persona{Name: name, LastName: lastName, Edad: edad}
}
