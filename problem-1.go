package main

import (
	"fmt"
)

// a Struct Matrix
// containing number of rows, number of columns, 2D slice
type Matrix struct{
	num_rows int
	num_cols int
	elements [][]int
}

// gets number of rows
func (matrix *Matrix) getNumRows() int{
	return matrix.num_rows
}

// get number of columns
func (matrix *Matrix) getNumCols() int{
	return matrix.num_cols
}

// sets value of element at x, y
func (matrix *Matrix) setElement(value, x, y int) {
	matrix.elements[x][y] = value
	return ;
}

// adds two matrix elements and return a new struct Matrix
func (matrix1 *Matrix) addMatrix(matrix2 *Matrix) Matrix{
	var num_rows = matrix1.getNumRows()
	var num_cols = matrix1.getNumCols()
	var res_matrix = Matrix{num_rows: num_rows, num_cols: num_cols}
	res_matrix.elements = make([][]int, num_rows)
	for row:=0; row<num_rows; row++ {
		res_matrix.elements[row] = make([]int, num_cols)
		for col:=0; col<num_cols; col++{
			res_matrix.elements[row][col] = matrix1.elements[row][col] + matrix2.elements[row][col]
		}
	}
	return res_matrix
}

func main(){

	var matrix1 = Matrix{num_rows: 3, num_cols: 3}
	matrix1.elements = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var matrix2 = Matrix{num_rows: 3, num_cols: 3}
	matrix2.elements = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix1.getNumRows())
	fmt.Println(matrix1.getNumCols())

	matrix1.setElement(10, 0, 0)

	fmt.Println(matrix1.addMatrix(&matrix2))

}
