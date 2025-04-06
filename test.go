package gopackagetesting

import "fmt"

type Input struct {
	Name string
	Age  int
}

func (p *Input) Print() {
	fmt.Println("Input:", p.Name)
}
