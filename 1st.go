package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	row int
	col int
	mat [][]int
}

func (m *Matrix) getnoRows() int {
	return m.row
}
func (m *Matrix) getnoCols() int {
	return m.col
}
func (m *Matrix) setPos(i, j, val int) {
	m.mat[i][j] = val
}
func (m *Matrix) printmat() {
	for _, x := range m.mat {
		for _, y := range x {
			fmt.Printf("%v ", y)
		}
		fmt.Printf("\n")
	}
}
func (m *Matrix) printmatJson() {
	data, _ := json.Marshal(m.mat)
	fmt.Println(string(data))
}
func (m *Matrix) init2d() {
	for i, _ := range m.mat {
		m.mat[i] = make([]int, m.col)
	}
}

func (a *Matrix) addMat(b *Matrix) {
	if a.row != b.row || a.col != b.col {
		fmt.Println("Incompatible")
		return
	}
	c := &Matrix{3, 4, make([][]int, 3)}
	c.init2d()
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			c.mat[i][j] = a.mat[i][j] + b.mat[i][j]
		}
	}
	c.printmatJson()
}

func main() {
	fmt.Println("hello world")
	m := &Matrix{3, 4, make([][]int, 3)}
	m.init2d()
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			m.setPos(i, j, i+j)
		}
	}
	fmt.Println(m.getnoRows())
	m.printmat()
	m.printmatJson()

	b := &Matrix{3, 4, make([][]int, 3)}
	b.init2d()
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			b.setPos(i, j, i-j)
		}
	}
	fmt.Println("print 2nd mat")
	b.printmatJson()
	fmt.Println("resultant matrix")
	m.addMat(b)

}
