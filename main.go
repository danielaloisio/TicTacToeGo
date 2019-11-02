package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var positions = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var player = "O"

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {

	var position int

	fmt.Printf("Digite uma posição para jogar \n")
	printBoard()

	for {

		var check int = checkStatus()
		if check == 1 {
			fmt.Printf("Player %s  won\n", player)
			break
		} else if check == 2 {
			fmt.Printf("End game, nobody won\n")
			break
		}

		fmt.Scanf("%d\n", &position)

		if position > 0 && position <= 9 {
			if positions[position] != "X" && positions[position] != "O" {

				positions[position] = nexPlayer()
			}
		}

		newClearScreen()
		printBoard()
	}
}

func nexPlayer() string {
	if player == "O" {
		player = "X"
	} else {
		player = "O"
	}

	return player
}

func newClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func printBoard() {

	fmt.Printf(" %s  |  %s  | %s \n", positions[1], positions[2], positions[3])
	fmt.Printf("---         --- \n")
	fmt.Printf(" %s  |  %s  | %s \n ", positions[4], positions[5], positions[6])
	fmt.Printf("---         --- \n")
	fmt.Printf(" %s  |  %s  | %s \n ", positions[7], positions[8], positions[9])
}

func checkStatus() int {
	//horizontal
	if positions[1] == positions[2] && positions[2] == positions[3] {
		return 1
	} else if positions[4] == positions[5] && positions[5] == positions[6] {
		return 1
	} else if positions[7] == positions[8] && positions[8] == positions[9] {
		return 1
	}

	//vertical
	if positions[1] == positions[4] && positions[4] == positions[7] {
		return 1
	} else if positions[2] == positions[5] && positions[5] == positions[8] {
		return 1
	} else if positions[3] == positions[6] && positions[6] == positions[9] {
		return 1
	}

	//diagonal
	if positions[1] == positions[5] && positions[5] == positions[9] {
		return 1
	} else if positions[3] == positions[5] && positions[5] == positions[7] {
		return 1
	}

	//check final without a winner
	if positions[1] != "1" && positions[2] != "2" && positions[3] != "3" && positions[4] != "4" && positions[5] != "5" && positions[6] != "6" && positions[7] != "7" && positions[8] != "8" && positions[9] != "9" {
		return 2
	}

	return -1
}
