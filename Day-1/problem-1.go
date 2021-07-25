package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Matrix containing number of rows, number of columns, 2D slice
type Matrix struct {
	NumRows  int     `json:"NumRows"`
	NumCols  int     `json:"NumCols"`
	Elements [][]int `json:"Elements"`
}

// gets number of rows
func (matrix *Matrix) getNumRows() int {
	return matrix.NumRows
}

// get number of columns
func (matrix *Matrix) getNumCols() int {
	return matrix.NumCols
}

func (matrix *Matrix) init() {
	numRows := matrix.getNumRows()
	numCols := matrix.getNumCols()
	matrix.Elements = make([][]int, numRows)
	val := 0
	for row := 0; row < numRows; row++ {
		matrix.Elements[row] = make([]int, numCols)
		for col := 0; col < numCols; col++ {
			val++
			matrix.setElementAt(row, col, val)
		}
	}
}

// sets value of element at row, col
func (matrix *Matrix) setElementAt(row, col, value int) error {
	if row < 0 || row >= matrix.getNumRows() || col < 0 || col >= matrix.getNumCols() {
		return errors.New("index out of bound")
	}
	matrix.Elements[row][col] = value
	return nil
}

func (matrix *Matrix) getElementAt(row, col int) (int, error) {
	if row < 0 || row >= matrix.getNumRows() || col < 0 || col >= matrix.getNumCols() {
		return 0, errors.New("index out of bound")
	}

	value := matrix.Elements[row][col]
	return value, nil
}

// addMatrix function adds two matrix elements and return a new struct Matrix
func addMatrix(matrix1, matrix2 *Matrix) *Matrix {
	var numRows = matrix1.getNumRows()
	var numCols = matrix1.getNumCols()
	var resMatrix = Matrix{NumRows: numRows, NumCols: numCols}
	resMatrix.init()
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			val1, er1 := matrix1.getElementAt(row, col)
			if er1 != nil {
				fmt.Println(er1)
			}
			val2, er2 := matrix2.getElementAt(row, col)
			if er2 != nil {
				fmt.Println(er2)
			}
			resMatrix.setElementAt(row, col, val1+val2)
		}
	}
	return &resMatrix
}

func (matrix *Matrix) printAsJson() {
	data, err := json.Marshal(matrix)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Matrix: %s\n", string(data))
}

func main() {

	var matrix1 = Matrix{NumRows: 3, NumCols: 3}
	matrix1.init()

	var matrix2 = Matrix{NumRows: 3, NumCols: 3}
	matrix2.init()

	err := matrix1.setElementAt(0, 0, 10)
	if err != nil {
		fmt.Println(err)
	}

	matrix1.printAsJson()
	matrix2.printAsJson()
	matrix3 := *addMatrix(&matrix1, &matrix2)
	matrix3.printAsJson()
}
