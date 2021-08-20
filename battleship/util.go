package main

import (
	"fmt"
	"math"
)

func printBoard(board [][]string) {
	letters := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 1; i < 11; i++ {
		fmt.Println()
		x := i
		fmt.Print(x, " ")
		if i != 10 {
			fmt.Print(" ")
		}
		for j := 0; j < 10; j++ {
			fmt.Print(board[i-1][j], " ")
		}
	}
	fmt.Println()
	fmt.Print("   ")
	for j := 0; j < 10; j++ {
		fmt.Print(letters[j], " ")
	}
	fmt.Println()
	fmt.Println()
}

func template() [][]string {
	board := [][]string{}
	for i := 0; i < 10; i++ {
		row := []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
		board = append(board, row)
	}
	return board
}

func isInvalid(sx, sy, fx, fy int) bool {
	if sx < 0 || sy < 0 || fx < 0 || fy < 0 || sx >= 10 || sy >= 10 || fx >= 10 || fy >= 10 {
		return true
	}
	return false
}

func isTaken(board [][]string, sx, sy, fx, fy, size int) bool {
	if sx != fx && sy != fy {
		return true
	}
	if sx == fx {
		var y int
		if sy < fy {
			y = sy
		} else {
			y = fy
		}
		for i := 0; i < size; i++ {
			if board[sx][y+i] != "0" {
				return true
			}
		}
	}
	if sy == fy {
		var x int
		if sx < fx {
			x = sx
		} else {
			x = fx
		}
		for i := 0; i < size; i++ {
			if board[x+i][sy] != "0" {
				return true
			}
		}
	}
	return false
}

func isWrongDistance(sx, sy, fx, fy, size int) bool {
	var dist float64
	if sx == fx {
		dist = float64(fy - sy)
		if math.Abs(dist) != float64(size-1) {
			return true
		} else {
			return false
		}
	} else if sy == fy {
		var dist float64 = float64(fx - sx)
		if math.Abs(dist) != float64(size-1) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func isDiagonal(sx, sy, fx, fy int) bool {
	if sx != fx && sy != fy {
		return true
	} else {
		return false
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getIndex(x, y int) int {
	return (x * 10) + y%10
}

func buildHealthMap(health map[string]int) {
	health["carrier"] = 5
	health["battleship"] = 4
	health["cruiser"] = 3
	health["sub"] = 3
	health["destroyer"] = 2
}
