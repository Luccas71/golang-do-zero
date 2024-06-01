package main

import "fmt"

func main() {

	fmt.Printf("Tipo: %T - valor: %v\n", true, true)
	fmt.Printf("Tipo: %T - valor: %v\n", 1.23, 1.23)
	fmt.Printf("Tipo: %T - valor: %v\n", 1, 1)
	fmt.Printf("Tipo: %T - valor: %v\n", "1", "1")
	fmt.Printf("Tipo: %T - valor: %v\n", "lucas", "lucas")
}

// Tipos
// bool => true/ false
// float => 64/32 ponto flutuante (12.2)
// int => (...,-1, 0, 23,...)
// string => sequencia de bytes
