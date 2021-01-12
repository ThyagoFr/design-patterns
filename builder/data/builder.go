package data

import "strings"

type PersonBuilder struct {
	person *Person
}

func (pb *PersonBuilder) Name(name string) *PersonBuilder {
	pb.person.name = name
	return pb
}

func (pb *PersonBuilder) Email(email string) *PersonBuilder {
	if !strings.Contains(email, "@") {
		panic("invalid email")
	}
	pb.person.email = email
	return pb
}

func (pb *PersonBuilder) Age(age int) *PersonBuilder {
	pb.person.age = age
	return pb
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}
