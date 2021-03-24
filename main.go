package main

import(
	// "bufio"
	// "strings"
	// "strconv"
	// "time"
	"fmt"
	"os"
	"os/exec"
)

func drawBoard (pieces [9]string, player int){
	fmt.Println("\nThis is how we do the Tic and the Tac with the Toe")
	fmt.Printf("It is player %d's turn \n", player)
	fmt.Println()
	for i, s := range pieces{
		i += 1
		if s == ""{
			fmt.Print(" ", i)
		} else{
			fmt.Print(" ", s)
		}

		if i%3 == 0 {
			fmt.Println("\n___ ___ ___")
		} else{
			fmt.Print(" |")
		}
	}
}

func getBox(pieces [9]string, player int)int{
	turn, boxInLimit, boxNotFilled := 0, false, false
	var index int
	boxInLimitErr := "The value given for the box is out of bounds.\n"
	boxNotFilledErr := "The box you provided was already filled in a previous move.\n"
	var errString string

	for !(boxInLimit && boxNotFilled) {
		drawBoard(pieces, player)
		if turn > 0{
			errString = ""
			// Decides which errors are applicaple for box entry
			if !boxInLimit {errString += boxInLimitErr} else if !boxNotFilled {errString += boxNotFilledErr}
		}

		fmt.Println("Which number would you like to change?")
		if errString != "" {fmt.Print(errString)} // Prints error
		fmt.Printf("Give your number here: ")

		fmt.Scanln(&index)
		// Next three clear the console after each question has been asked
		c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
		turn++
		if index != 0 && index >= 1 && index <= 9{
			boxInLimit = true
			if pieces[index - 1] == ""{
				boxNotFilled = true
			}
		}	else{	boxInLimit = false}
	}
	return index
}

func checkWinner(pieces [9]string)(bool, int){
	isWinner := false
	winner := 0

	for i := 0; i <= 8; i+=1{
		if pieces[i] != ""{
			if i <= 2{
				if (pieces[i] == pieces[i+3] && pieces[i] == pieces[i+6]){isWinner = true}
				if i == 0 && (pieces[0] == pieces [4] && pieces[0] == pieces[8]) {isWinner = true}
				if i == 2 && (pieces[2] == pieces [4] && pieces[2] == pieces[6]) {isWinner = true}
			} else if i%3 == 0 && (pieces[i] == pieces[i+1] && pieces[i] == pieces[i+2]){
				isWinner = true
			}
			if isWinner{
				if pieces[i] == "X" {winner = 1} else {winner = 2}
				break
			}
		}
	}
	return isWinner, winner
}

func computerMove(pieces [9]string)int{
	var move int
	
	return move
}

func main (){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Println("\nWelcome To Tic Tac Toe")
	var pieces [9]string
	turn := 0
	player, winningPlayer := 1, 0
	gameOver := false

	for !gameOver {
		index := getBox(pieces, player)
		if turn%2 == 0{
			pieces[index - 1] = "X"
			player = 2
		} else {
			pieces[index - 1] = "O"
			player = 1
		}
		turn ++
		gameOver, winningPlayer = checkWinner(pieces)
	}

	fmt.Println("Game Over. \nPlayer", winningPlayer, "is the winner. \nPress Enter to exit.")
	fmt.Scanln()
}