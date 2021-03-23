package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main (){
	fmt.Println("Welcome To Tic Tac Toe")

	var pieces [9]string

	for i, s := range pieces{
		i += 1
		// fmt.Println(i, s)
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
	// for i := 1; i <= 3; i++{
	// 	for j:= 1; j <= 3;j++{
	// 		fmt.Print(pieces[j+i], "|")
	// 	}
	// 	fmt.Println()
	// }

	//var String
	var stringValue string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Which number would you like to change? \n")
	//user input
	stringValue, _ = reader.ReadString('\n')
	stringValue = strings.TrimSuffix(stringValue, "\n")
	// originalValue = stringValue
	//remove spaces
	stringValue = strings.ReplaceAll(stringValue, " ", "")
	//set string to lower case
	// stringValue = strings.ToLower(stringValue)
	// //if contains substring "i", "a" and "n"
	// if strings.Contains(stringValue, "i") && strings.Contains(stringValue, "a") && strings.Contains(stringValue, "n") {
	// 		// if "i" is a prefix and "n" a suffix
	// 		firstChar := strings.HasPrefix(stringValue, "i")
	// 		lastChar := strings.HasSuffix(stringValue, "n")
	// 		switch {
	// 		//if is a prefix and suffix so Found
	// 		case firstChar && lastChar:
	// 				fmt.Printf("Found in %s", originalValue)
	// 		//if isnt a prefix and suffix so Not Found
	// 		default:
	// 				fmt.Printf("Not Found in %s", originalValue)
	// 		}
	// 		//if there's no "i", "a" and "n"
	// } else {
	// 		fmt.Printf("Not Found in %s", originalValue)
	// }
	index, _ := strconv.Atoi(stringValue)

	pieces[index - 1] = "H"

	for i, s := range pieces{
		i += 1
		// fmt.Println(i, s)
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