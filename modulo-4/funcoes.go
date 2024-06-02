package main

import "fmt"

func main() {
	fmt.Println(soma(158, 256))

	nome1, nome2 := printaNome("Lucas")
	fmt.Println(nome1)
	fmt.Println(nome2)
}

func soma(x, y int) int {
	return x + y
}

func printaNome(nome string) (string, string) {
	return nome, nome
}
