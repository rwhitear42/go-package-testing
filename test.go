package gopackagetesting

import "fmt"

type Input struct {
	Name     string
	Age      int
	Shoesize float32
}

func (p *Input) Print() {
	fmt.Println("Input:", p.Name, p.Age, p.Shoesize)
}
