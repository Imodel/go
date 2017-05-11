package martix

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {
	p := new(person)
	p.name = name
	p.age = age
	return p
}
