package main

import (
	"fmt"

	pattern "github.com/thyagofr/design-patterns/builder/pattern"
)

// Builder Pattern -
// Utilizado quando se tem um objeto complicado de se construir/instanciar

func main() {

	builder := pattern.NewPersonBuilder()
	builder.
		Email("thyagofr@alu.ufc.br").
		Name("Thyago Freitas da Silva").
		Age(22)
	person := *builder.Build()
	fmt.Println(person)
}
