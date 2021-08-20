package main

import (
	"fmt"
	"strconv"
)

var myCoords = make(map[int]string)

func buildMyBoard() [][]string {
	board := template()
	printBoard(board)
	fmt.Println("Place your carrier (5 holes)")
	placeMyBoat(board, 5, "carrier")
	printBoard(board)
	fmt.Println("Place your battleship (4 holes)")
	placeMyBoat(board, 4, "battleship")
	printBoard(board)
	fmt.Println("Place your cruiser (3 holes)")
	placeMyBoat(board, 3, "cruiser")
	printBoard(board)
	fmt.Println("Place your submarine (3 holes)")
	placeMyBoat(board, 3, "sub")
	printBoard(board)
	fmt.Println("Place your destroyer (2 holes)")
	placeMyBoat(board, 2, "destroyer")
	return board
}

func myMove(CPUBoard [][]string, visual [][]string) {
	var guess string
	for {
		fmt.Println("Pick a coordinate")
		fmt.Scanf("%s", &guess)
		x, y := parseCoordinate(guess)
		if CPUBoard[x][y] == "x" || CPUBoard[x][y] == "H" {
			fmt.Println("Coordinate already struck")
		} else if CPUBoard[x][y] == "B" {
			println("Hit!")
			CPUBoard[x][y] = "H"
			visual[x][y] = "H"
			CPUHealth -= 1
			s := CPUCoords[getIndex(x, y)]
			CPUHealthMap[s] -= 1
			if CPUHealthMap[s] == 0 {
				fmt.Printf("%s sunk!\n", s)
			}
			break
		} else {
			println("Miss!")
			CPUBoard[x][y] = "x"
			visual[x][y] = "x"
			break
		}
	}
}

func placeMyBoat(board [][]string, size int, boat string) {
	var start, finish string
	for {
		fmt.Println("Give starting coordinate")
		fmt.Scanf("%s", &start)
		sx, sy := parseCoordinate(start)
		fmt.Println("Give ending coordinate")
		fmt.Scanf("%s", &finish)
		fx, fy := parseCoordinate(finish)
		if isInvalid(sx, sy, fx, fy) {
			fmt.Println("Coordinates must be in range of board")
		} else if isDiagonal(sx, sy, fx, fy) {
			fmt.Println("Ships cannot be placed diagonally")
		} else if isWrongDistance(sx, sy, fx, fy, size) {
			fmt.Println("Coordinate span does not match ship size")
		} else if isTaken(board, sx, sy, fx, fy, size) {
			fmt.Println("Spots already taken")
		} else {
			fillMySpots(board, sx, sy, fx, fy, size, boat)
			break
		}
	}
	return
}

func fillMySpots(board [][]string, sx, sy, fx, fy, size int, boat string) {
	if sx == fx {
		y := min(sy, fy)
		for i := 0; i < size; i++ {
			board[sx][y+i] = "B"
			index := getIndex(sx, y+1)
			myCoords[index] = boat
		}
		return
	} else {
		x := min(sx, fx)
		for i := 0; i < size; i++ {
			board[x+i][sy] = "B"
			index := getIndex(x+i, sy)
			myCoords[index] = boat
		}
		return
	}
}

func parseCoordinate(s string) (int, int) {
	if len(s) < 2 || len(s) > 3 {
		fmt.Println("Invalid coordinate")
		return -1, -1
	}
	x := int(s[0]) - 65
	if x >= 10 {
		fmt.Println("Invalid coordinate")
		return -1, -1
	}
	runes := []rune(s)
	str := string(runes[1:])
	z, err := strconv.ParseInt(str, 10, 64)
	if err != nil || z < 0 || z > 10 {
		fmt.Println("Invalid coordinate")
		return -1, -1
	}
	y := int(z) - 1
	return y, x
}
