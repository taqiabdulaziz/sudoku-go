package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var board [][]int

func solve(board [][]int) [][]int {
	var keyValue map[string]int
	keyValue = map[string]int{}
	var empty []map[string]int

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				keyValue["row"] = i
				keyValue["col"] = j
				keyValue["val"] = 0

				empty = append(empty, keyValue)
				keyValue = map[string]int{}
			}
		}
	}

	for i, v := range empty {
		row := v["row"]
		col := v["col"]
		val := v["val"]

		for {
			if val > 9 {
				board[row][col] = 0
				empty[i]["val"] = 0
				i -= 2
				break
			}
			if checkVertical(col, val) && checkHorizontal(board[row], val) {
				board[row][col] = val
				empty[i]["val"] = val
				break
			} else {
				val++
			}
		}
	}

	return board
}

func checkVertical(col, val int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if j == col && board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func checkHorizontal(horizontalLine []int, val int) bool {
	for i := 0; i < len(horizontalLine); i++ {
		if horizontalLine[i] == val {
			return false
		}
	}
	return true
}

func generateMultiDimensionalSlice(uintArr []uint8) [][]int {
	var responseArr [][]int
	responseArr = [][]int{}
	var counter int

	for i := 0; i < 9; i++ {
		var arrTemp []int
		for j := 0; j < 9; j++ {
			parsedInt, _ := strconv.Atoi(string(uintArr[counter]))
			arrTemp = append(arrTemp, parsedInt)
			counter++
		}
		responseArr = append(responseArr, arrTemp)
		arrTemp = nil
	}
	return responseArr
}

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("err on", err)
	}

	parsedFile, err := ioutil.ReadAll(file)

	board = generateMultiDimensionalSlice(parsedFile)
	var solvedBoard [][]int
	solvedBoard = solve(board)
	for i := 0; i < len(solvedBoard); i++ {
		fmt.Println(solvedBoard[i])
	}

}
