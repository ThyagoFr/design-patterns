package main

import (
	"fmt"

	"github.com/thyagofr/design-patterns/builder/data"
)

// Builder Pattern -
// Utilizando quando se tem um objeto complicado de se construir/instanciar

func main() {

	builder := data.NewPersonBuilder()
	builder.
		Email("thyagofr@alu.ufc.br").
		Name("Thyago Freitas da Silva").
		Age(22)
	person := *builder.Build()
	fmt.Println(person)
}
