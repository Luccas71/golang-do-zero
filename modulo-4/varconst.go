package main

import "fmt"

func main() {

	//variaveis => seu valor pode mudar, mas NUNCA o seu tipo
	var nome string
	nome = "Lucas"
	fmt.Println(nome)

	//outras formas de declarar uma variavel

	//inferencia de tipo
	var a = "Maciel"
	fmt.Println(a)

	//multiplas variaveis de mesmo tipo
	var b, c int = 12, 13
	fmt.Println(b)
	fmt.Println(c)

	//declaração e atribuição (shorthand)
	d := true
	fmt.Println(d)

	//Constantes
	//valor IMUTÁVEL
	const PI = 3.14
}
