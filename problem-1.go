package main

import (
	"fmt"
)

// Matrix containing number of rows, number of columns, 2D slice
type Matrix struct{
	numRows  int
	numCols  int
	elements [][]int
}

// gets number of rows
func (matrix *Matrix) getNumRows() int{
	return matrix.numRows
}

// get number of columns
func (matrix *Matrix) getNumCols() int{
	return matrix.numCols
}

// sets value of element at x, y
func (matrix *Matrix) setElement(value, x, y int) {
	matrix.elements[x][y] = value
	return ;
}

// adds two matrix elements and return a new struct Matrix
func (matrix1 *Matrix) addMatrix(matrix2 *Matrix) Matrix{
	var numRows = matrix1.getNumRows()
	var numCols = matrix1.getNumCols()
	var resMatrix = Matrix{numRows: numRows, numCols: numCols}
	resMatrix.elements = make([][]int, numRows)
	for row:=0; row< numRows; row++ {
		resMatrix.elements[row] = make([]int, numCols)
		for col:=0; col< numCols; col++{
			resMatrix.elements[row][col] = matrix1.elements[row][col] + matrix2.elements[row][col]
		}
	}
	return resMatrix
}

func main(){

	var matrix1 = Matrix{numRows: 3, numCols: 3}
	matrix1.elements = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var matrix2 = Matrix{numRows: 3, numCols: 3}
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
