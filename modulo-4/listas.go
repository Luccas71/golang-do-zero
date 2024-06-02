package main

import "fmt"

func main() {

	//ARRAYS
	//declarando um array => var nomeDoArray[tamanho]tipo => var numeros[2]int
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[1])

	numInt := [6]int{1, 0, 23, 56, 5, 7}
	fmt.Println(numInt[2])

	//pega dos valores da pos[1] inclusive ate o valor anterior a pos[3] => [0, 23]
	fmt.Println(numInt[1:3])

	//pega os valores até a pos[2]
	fmt.Println(numInt[:2])

	//pega os valores da pos[3] inclusive e os valores adiante
	fmt.Println(numInt[3:])

	//SLICES
	//criando um slice vazio
	slice := make([]int, 2)
	slice[0] = 1
	slice[1] = 2
	fmt.Println(slice)

	//pegando tamanho do slice
	fmt.Println(len(slice))

	//adicionando items ao slice
	slice = append(slice, 5, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(len(slice))

}

/*
Listas

1- Arrays e Slices => todos elementos são do mesmo tipo
[1, 2, 4] => int
[2.2, 3.3, 4.04] => float
["Lucas", "Maciel"] => strings

2- Maps => pode ter elementos de vários tipos
	- estrutura chave-valor
	- [key] = value
	- map[string]int
		- { "lucas": 30, "mari": 24}
*/

/*
Array:
	- tamanho fixo
	- valor acessado por indice: array[1], array[3]
	- func len() retorna o tamanho do array
	- pouco usado devido ao tamanho fixo

Slice:
	- parecido com array, mas sem o tamanho fixo
	- valor acessado por indice: array[1], array[3]
	- func len() retorna o tamanho do array
	- func append() utilizada para adicionar valores
*/
