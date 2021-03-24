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

func checkWinner(pieces [9]string)bool{
	for i := 0; i <= 2; i+=1{
		if pieces[i] == pieces[i+3] && pieces[i] == pieces[i+6]{
			return true
		}
	}
	for i := 0; i <= 2; i+=3{
		if pieces[i] == pieces[i+1] && pieces[i] == pieces[i+2]{
			return true
		}
	}
	return false
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
			if !boxInLimit {errString += boxInLimitErr} else if !boxNotFilled {errString += boxNotFilledErr}
		}

		fmt.Println("Which number would you like to change?")
		if errString != "" {fmt.Print(errString)}
		fmt.Printf("Give your number here: ")

		fmt.Scanln(&index)
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

func main (){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
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