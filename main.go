package main

import(
	// "bufio"
	"fmt"
	"os"
	"os/exec"
	// "strings"
	// "strconv"
	// "time"
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
	boxInLimit := false
	var index int

	for !boxInLimit {
		drawBoard(pieces, player)
		fmt.Println("Which number would you like to change?")
		fmt.Printf("Give your number here: ")

		fmt.Scanln(&index)
		c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
		if index != 0 && index >= 1 && index <= 9{
			boxInLimit = true
		}
	}
	return index
}

func main (){
	fmt.Println("\nWelcome To Tic Tac Toe")
	var pieces [9]string
	turn := 0
	player := 1

	for true {
		index := getBox(pieces, player)
		if turn%2 == 0{
			pieces[index - 1] = "X"
			player = 2
		} else {
			pieces[index - 1] = "O"
			player = 1
		}
		turn ++
	}
}