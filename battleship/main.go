package main

var myHealth = 17
var CPUHealth = 17
var myHealthMap = make(map[string]int)
var CPUHealthMap = make(map[string]int)

func main() {
	println("Let's play Battleship!")
	myBoard := buildMyBoard()
	printBoard(myBoard)
	println("Enemy building board...")
	CPUBoard := buildCPUBoard()
	printBoard(CPUBoard)
	playGame(myBoard, CPUBoard)
}

func playGame(myBoard, CPUBoard [][]string) {
	visual := template()
	buildHealthMap(myHealthMap)
	buildHealthMap(CPUHealthMap)
	for {
		println()
		println("Your board")
		printBoard(myBoard)
		println("Enemy board")
		printBoard(visual)
		myMove(CPUBoard, visual)
		if CPUHealth == 0 {
			println("You win!")
			break
		}
		println("Your enemy's turn")
		CPUMove(myBoard)
		if myHealth == 0 {
			println("You lose!")
			break
		}
	}
}
