package human

import "fmt"

type Human interface {
	GetName() string
	Walk(steps int)
	Talk(words int)
}

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) Walk(steps int) {
	fmt.Printf("%s walked %d steps\n", p.Name, steps)
}

func (p *Person) Talk(words int) {
	fmt.Printf("%s said %d words\n", p.Name, words)
}
