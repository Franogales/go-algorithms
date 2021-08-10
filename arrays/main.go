package main

import "fmt"

/*
existen 2 tipos de arrays , estaticos y dinamicos
los estaticos son los que hacen uso de las listas
que basicamente es agarrar un array existente y copiarlo y moverlo
completamente a otro espacio de memoria completamente diferente.
En Go existen dos palabras para manejar los elementos en memoria
array y slice, donde array es un array estatico y slice es un array dinamico.
La ventaja del Slice es clara, ya que no te preocupas por la cantidad de elementos en el array
pero la ventaja del Array es que puedes manejar mejor la memoria
*/

var index = 4

func main() {
	primes := [6]int{2, 3, 5, 7}
	push(12, primes[:])
	fmt.Println(primes)
	push(15, primes[:])
	fmt.Println(primes)
	pop(primes[:])
	fmt.Println(primes)
	pop(primes[:])
	fmt.Println(primes)
}

func push(element int, array []int) {
	array[index] = element
	index++
}
func pop(array []int) {
	index--
	array[index] = 0
}

// func add(array []int) {
// 	newSize := len(array) + 1
// 	get
// }
