package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Matrix struct {
	numOfRows, numOfColumns int
	twoDMatrix [][]int
}
func Initialiser(rows, cols int) *Matrix{
	matrix := &Matrix{
		numOfRows:  rows,
		numOfColumns: cols,
		twoDMatrix: make([][]int, rows) , //Creates a regular array instead of 2D
	}
	for iterator:=range matrix.twoDMatrix{
		matrix.twoDMatrix[iterator] = make([]int, cols)
	}
	return matrix
}
func (ourMatrix Matrix) row_size() int {
	return ourMatrix.numOfRows
}
func (ourMatrix Matrix) column_size() int{
	return ourMatrix.numOfColumns
}

func main(){
	ourMatrix := Initialiser(3,3)
	fmt.Println(ourMatrix.row_size())
	fmt.Println(ourMatrix.column_size())
	for row:=0 ; row < ourMatrix.row_size() ; row++{
		for col:=0; col < ourMatrix.column_size(); col++{
			ourMatrix.twoDMatrix[row][col] = row + col
		}
	}
	fmt.Println(ourMatrix)
	secondMatrix := Initialiser(3,3)
	for row:=0 ; row < secondMatrix.row_size() ; row++{
		for col:=0; col < secondMatrix.column_size(); col++{
			secondMatrix.twoDMatrix[row][col] = row + 2*col
		}
	}
	sumMatrix := Initialiser(3,3)
	for row:=0 ; row < sumMatrix.row_size() ; row++{
		for col:=0; col < sumMatrix.column_size(); col++{
			sumMatrix.twoDMatrix[row][col] = ourMatrix.twoDMatrix[row][col] + secondMatrix.twoDMatrix[row][col]
		}
	}
	fmt.Println(sumMatrix)
	data,err := json.Marshal(ourMatrix)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("s\n",data)



}