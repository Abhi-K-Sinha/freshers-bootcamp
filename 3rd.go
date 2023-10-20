package main

import (
	"fmt"
)

type calculate interface {
	getSalary() float64
}
type fte struct{}
type contractor struct {
	days int
}
type freelance struct {
	hours int
}

func (m *fte) getSalary() float64 {
	return float64(15000)
}
func (m *contractor) getSalary() float64 {
	return float64(m.days * 100)
}
func (m *freelance) getSalary() float64 {
	if m.hours < 20 {
		return 0
	}
	return float64(m.hours * 100)
}

func main() {
	full := fte{}
	con := contractor{20}
	free := freelance{100}
	fmt.Print("fte ", full.getSalary(), "\n")
	fmt.Print("contractor ", con.getSalary(), "\n")
	fmt.Print("freelancer ", free.getSalary())
}
