package main

import (
	"fmt"
)

/*
dado un array de enteros, returna los indices de 2 numeros que se agregan a un objetivo dado.
Vamos a recibir dos argumentos
1 - array de numeros
2 - target (objetivo)
la suma de 2 numeros en el array deben de dar el valor del objetivo
por ejemplo si en el array contiene lo siguiente [1,3,7,9,2] y el objetivo es 11 = la suma de 9 y 2 = 11, entonces el programa debe retornar esos indices.
en ese caso 3 y 4 = [3,4]

no hay duplicados
todos los numeros son positivos (no numeros en negativos)
el array tiene mas de un elemento
el target debe de ser alcanzable por los elementos del array
solamente un par puede dar el target
*/

func main() {
	fmt.Println(getIndexWithHash([]int{1, 3, 7, 9, 2}, 11))
	fmt.Println(getIndex1([]int{1, 3, 7, 9, 2}, 11))
}

/*
la funcion getIndex1 tiene
tiempo = O(n^2)
espacio = O(1)
*/
func getIndex1(array []int, target int) []int {

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
la funcion getIndexWithHash tiene
tiempo = O(n)
espacio = O(n)
*/
func getIndexWithHash(array []int, target int) []int {
	// inicio el map vacio

	mapa1 := make(map[int]int)
	for i := 0; i < len(array); i++ {

		valueInMap := mapa1[array[i]]
		// si el valor no se encuetra en el map, returna 0
		if valueInMap > 0 {
			// si no se encontro el num
			return []int{valueInMap, i}
		} else {
			numToFind1 := target - array[i]
			mapa1[numToFind1] = i
		}
	}
	return nil
}
