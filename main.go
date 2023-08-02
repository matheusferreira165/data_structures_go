package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Sexo      string
}

func main() {
	pessoas := []Pessoa{
		{"Matheus", "Ferreira", 24, "Masculino"},
		{"Maria", "Ferreira", 24, "Feminino"},
		{"Joao", "Ferreira", 24, "Masculino"},
		{"Tereza", "Ferreira", 24, "Feminino"},
	}

	table := HashTable{}

	keys := make([]int, len(pessoas))

	for i, pessoa := range pessoas {
		keys[i] = table.Put(pessoa)
	}

	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Nome, p.Sobrenome, p.Idade, p.Sexo)
		}
	}

	table.Remove("Maria")

	tiago := table.Search("Tereza")
	fmt.Println(tiago)
}
