package main

import (
	"encoding/json"
	"fmt"
	"log"
	"error"
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
//to find out row size
func (ourMatrix Matrix) rowSize() int {
	return ourMatrix.numOfRows
}
//to find out column size
func (ourMatrix Matrix) columnSize() int{
	return ourMatrix.numOfColumns
}
// to set element
func (ourMatrix *Matrix) setElements(rowIndex, colIndex, Value int){
	if rowIndex >= ourMatrix.rowSize() || colIndex >= ourMatrix.columnSize() {
		fmt.Print("Can't work with these indices")
		return
	}
	 ourMatrix.twoDMatrix[rowIndex][colIndex] = Value
}
// to add 2 matrices
func (ourMatrix Matrix) addMatrix(secondMatrix , sumMatrix *Matrix) {
	for row:=0 ; row < sumMatrix.rowSize() ; row++ {
		for col := 0; col < sumMatrix.columnSize(); col++ {
			sumMatrix.twoDMatrix[row][col] = ourMatrix.twoDMatrix[row][col] + secondMatrix.twoDMatrix[row][col]
		}
	}
}
//to print matrix structure as json
func (ourMatrix Matrix) printAsJson(){
	data,err := json.Marshal(ourMatrix)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("s\n",data)
}
func main(){
	ourMatrix := Initialiser(3,3)
	fmt.Println(ourMatrix.rowSize())//a part
	fmt.Println(ourMatrix.columnSize())//b part
	for row:=0 ; row < ourMatrix.rowSize() ; row++{
		for col:=0; col < ourMatrix.columnSize(); col++{
			ourMatrix.twoDMatrix[row][col] = row + col
		}
	}
	fmt.Println(ourMatrix.twoDMatrix)
	ourMatrix.setElements(2,2,10)//c part
	fmt.Println(ourMatrix.twoDMatrix)
	//initialising the second matrix
	secondMatrix := Initialiser(3,3)
	for row:=0 ; row < secondMatrix.rowSize() ; row++{
		for col:=0; col < secondMatrix.columnSize(); col++{
			secondMatrix.twoDMatrix[row][col] = row + 2*col
		}
	}
	sumMatrix := Initialiser(3,3)//matrix that will store the sum of 2 matrices
	ourMatrix.addMatrix(secondMatrix,sumMatrix)//d part
	fmt.Println(sumMatrix.twoDMatrix)
	ourMatrix.printAsJson()//e part



}