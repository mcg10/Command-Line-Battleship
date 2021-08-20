package main

import (
	"fmt"
	"math/rand"
	"time"
)

type square struct {
	x         int
	y         int
	direction string
}

var direction int = 2
var queue []square
var CPUCoords = make(map[int]string)

func buildCPUBoard() [][]string {
	board := template()
	placeCPUBoat(board, 5, "carrier")
	placeCPUBoat(board, 4, "battleship")
	placeCPUBoat(board, 3, "cruiser")
	placeCPUBoat(board, 3, "sub")
	placeCPUBoat(board, 2, "destroyer")
	return board
}

func CPUMove(myBoard [][]string) {
	vertical := 0
	horizontal := 1
	for {
		var x, y int
		var point square
		var flip bool
		if len(queue) != 0 {
			flip = true
			point = queue[0]
			queue = queue[1:]
			x, y = point.x, point.y
		} else {
			flip = false
			direction = 2
			rand.Seed(time.Now().UnixNano())
			x = rand.Intn(10)
			y = rand.Intn(10)
		}
		if isInvalid(x, y, 0, 0) {

		} else {
			if myBoard[x][y] == "❌" || myBoard[x][y] == "🔥" {

			} else if myBoard[x][y] == "B" {
				println("Hit!")
				myBoard[x][y] = "🔥"
				myHealth -= 1
				if flip {
					if point.direction == "vertical" {
						direction = vertical
					} else {
						direction = horizontal
					}
				}
				if direction == 2 {
					queue = append(queue, square{x + 1, y, "vertical"})
					queue = append(queue, square{x - 1, y, "vertical"})
					queue = append(queue, square{x, y + 1, "horizontal"})
					queue = append(queue, square{x, y - 1, "horizontal"})
				} else if direction == vertical {
					queue = append(queue, square{x + 1, y, "vertical"})
					queue = append(queue, square{x - 1, y, "vertical"})
				} else if direction == horizontal {
					queue = append(queue, square{x, y + 1, "horizontal"})
					queue = append(queue, square{x, y - 1, "horizontal"})
				}
				s := myCoords[getIndex(x, y)]
				myHealthMap[s] -= 1
				if myHealthMap[s] == 0 {
					fmt.Printf("%s sunk!\n", s)
					direction = 2
				}
				break
			} else {
				println("Miss!")
				myBoard[x][y] = "❌"
				break
			}
		}
	}
}

func placeCPUBoat(board [][]string, size int, boat string) {
	vertical := 0
	for {
		rand.Seed(time.Now().UnixNano())
		var fx, fy int
		sx := rand.Intn(10)
		sy := rand.Intn(10)
		placement := rand.Intn(2)
		if placement == vertical {
			fy = sy
			if sx+size-1 >= 10 {
				fx = sx - size + 1
			} else {
				fx = sx + size - 1
			}
		} else {
			fx = sx
			if fy+size-1 >= 10 {
				fy = sy - size + 1
			} else {
				fy = sy + size - 1
			}
		}
		if isInvalid(sx, sy, fx, fy) {
		} else if isDiagonal(sx, sy, fx, fy) {
		} else if isWrongDistance(sx, sy, fx, fy, size) {
		} else if isTaken(board, sx, sy, fx, fy, size) {
		} else {
			fillCPUSpots(board, sx, sy, fx, fy, size, boat)
			break
		}
	}
	return
}

func fillCPUSpots(board [][]string, sx, sy, fx, fy, size int, boat string) {
	if sx == fx {
		y := min(sy, fy)
		for i := 0; i < size; i++ {
			board[sx][y+i] = "🛥️"
			index := getIndex(sx, y+i)
			CPUCoords[index] = boat
		}
		return
	} else {
		x := min(sx, fx)
		for i := 0; i < size; i++ {
			board[x+i][sy] = "🛥️"
			index := getIndex(x+i, sy)
			CPUCoords[index] = boat
		}
		return
	}
}
