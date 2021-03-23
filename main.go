package main

import(
	// "bufio"
	"fmt"
	"os"
	"os/exec"
	// "strings"
	// "strconv"
)

func drawBoard (pieces [9]string){
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

func getBox(pieces [9]string)int{
	correctInput := true
	var index int
	drawBoard(pieces)

	for correctInput {
		fmt.Println("Which number would you like to change?")
		fmt.Print("Give your number here: ")
		fmt.Scanln("%d", &index)
		fmt.Println(index)
		if index != 0 && index >= 1 && index <= 9{
			break
		}
	}
	return index
}

func main (){
	var pieces [9]string
	turn := 0

	for true {
		fmt.Println("\nWelcome To Tic Tac Toe")
		index := getBox(pieces)
		if turn%2 == 0{
			pieces[index - 1] = "X"
		} else {
			pieces[index - 1] = "O"
		}
		turn ++
		c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
	}
}