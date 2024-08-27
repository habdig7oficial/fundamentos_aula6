/*
Arquivo Go
Licença Apache 2.0
Copyright 2024 Mateus Felipe da Silveira Vieira
*/
package main

import (
	"fmt"
)

type Funcionario struct {
	nome string
	salario float32
}

func main() {
	funcionarios := []Funcionario {
		Funcionario{
           nome: "Mateus",
		   salario: 100.00,
	    },
		Funcionario{
			nome: "Marcos",
			salario: 600.50,
		},
    }

	var novo_salario []Funcionario = funcionarios
	
	for i := 0; i < len(funcionarios); i++ {
		if funcionarios[i].salario <= 500 {
			novo_salario[i].salario = funcionarios[i].salario + (funcionarios[i].salario / 100) * 20
		} else {
           novo_salario[i].salario = funcionarios[i].salario + (funcionarios[i].salario / 100) * 10
		}
	}

	fmt.Println(novo_salario)
}