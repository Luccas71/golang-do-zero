package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

type Active struct {
	Cliente  Cliente
	IsActive bool
}

func main() {

	//formas de instaciação
	fmt.Println(Cliente{"Lucas", 30})

	//forma mais indicada, passando chave e valor
	fmt.Println(Cliente{Nome: "Ademir", Idade: 52})

	//pode ser criada sem definir os elementos
	fmt.Println(Cliente{})

	//criando uma struct e atribuindo seu valor a uma variavel
	c1 := Cliente{Nome: "Joao", Idade: 23}

	//acessando os valores da struct
	fmt.Println(c1.Nome)

	//alterando valores
	c1.Nome = "Alfredo"
	fmt.Println(c1.Nome)

	//criando uma lista de pessoas
	clientes := []Cliente{
		{"Adao", 23},
		{"Gomes", 22},
	}
	fmt.Println(clientes)

	//map + struct
	clientes2 := map[string][]Cliente{
		"Programação": clientes,
		"Eng":         {{"Carlos", 19}, {"Lima", 22}},
	}
	fmt.Println(clientes2)

	//herdando campos de outra struct
	cliente3 := Active{Cliente{"Jose", 23}, true}
	fmt.Println(cliente3)
	fmt.Println(cliente3.Cliente.Nome)

	cliente4 := Active{
		Cliente{
			Nome:  "Valeria",
			Idade: 36,
		}, true,
	}
	fmt.Println(cliente4)
}

/*

	STRUCT
	- forma de criar a minha própria estrutura de dados
	- personalizar de acordo com a necessidade
	- aceita varios de tipos


*/
