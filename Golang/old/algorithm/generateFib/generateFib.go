// package generateFib
package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	matrix [][]int
	row    int // 行
	column int // 列
}

func main() {
	// fmt.Println(fib(10))
	// fmt.Println(quickPow(2, 8))
	a := generateStandardMatrix()
	// b := generateStandardMatrix()
	// res, _ := MatrixMultiply(a, b)
	// res, _ = MatrixMultiply(res, res)
	// res.printMatrix()
	c := a.quickMatrixPow(4)
	c.printMatrix()
	// fmt.Println(c.matrix[0][0])

}

func (mt *Matrix) Init(row, column int) error {
	if row < 0 || column < 0 {
		return errors.New("Argument error")
	}
	mt.row = row
	mt.column = column
	for i := 0; i < mt.column; i++ {
		tmp := make([]int, mt.row)
		mt.matrix = append(mt.matrix, tmp)
	}
	return nil
}

func generateStandardMatrix() *Matrix {
	tmp := new(Matrix)
	tmp.row = 2
	tmp.column = 2
	tmp.matrix = [][]int{
		{1, 1},
		{1, 0},
	}
	return tmp
}

func MatrixMultiply(a, b *Matrix) (*Matrix, error) {
	if a.column != b.row {
		return nil, errors.New("Argument error")
	}
	tmp := new(Matrix)
	tmp.Init(a.row, b.column)
	for i := 0; i < a.row; i++ {
		for z := 0; z < b.column; z++ {
			res := 0
			for q := 0; q < a.column; q++ {
				res = res + a.matrix[i][q]*b.matrix[q][z]
			}
			// tmp.matrix[i] = append(tmp.matrix[i], res)
			tmp.matrix[i][z] = res
		}
	}
	return tmp, nil
}

func (mt *Matrix) printMatrix() {
	if mt.column != 0 || mt.row != 0 {
		for i := 0; i < mt.row; i++ {
			for subI := 0; subI < mt.column; subI++ {
				fmt.Printf("%d ", mt.matrix[i][subI])
			}
			fmt.Println()
		}
	}
	return
}

func (mt *Matrix) Unit() {
	if mt.column < 0 && mt.row < 0 {
		return
	}
	for i := 0; i < mt.row; i++ {
		for z := 0; z < mt.column; z++ {
			if z == i {
				mt.matrix[i][z] = 1
			}
		}
	}
}

func quickPow(a, b int) int {
	res := 1
	if b == 0 {
		return res
	}
	if b&1 == 1 {
		res *= a
		b--
	}
	b >>= 1
	res *= quickPow(a, b) * quickPow(a, b)

	return res
}

func (mt *Matrix) quickMatrixPow(step int) *Matrix {
	var res *Matrix = new(Matrix)
	c := mt
	res.Init(2, 2)
	res.Unit()
	if step == 0 {
		return res
	}
	if step&1 == 1 {
		res, _ = MatrixMultiply(res, mt)
	}
	step >>= 1
	tmp, _ := MatrixMultiply(c.quickMatrixPow(step), c.quickMatrixPow(step))
	res, _ = MatrixMultiply(res, tmp)
	return res
}
