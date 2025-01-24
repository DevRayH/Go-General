package parent

import "fmt"

type Parent struct{
	Name string
}

func (p *Parent) Speak(){
	fmt.Println("Hello, my name is", p.Name)
}