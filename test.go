package gopackagetesting

import "fmt"

type Input struct {
	Name string
}

func (p *Input) Print() {
	fmt.Println("Input:", p)
}
