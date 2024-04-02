package main

import (
	"bufio"
	"fmt"
	"os"
)


func main ()  {

	welcom := "welcome to my go user input app"

	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the rating for our peezzer:")

	//comma error sntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating us", input)
	fmt.Printf("Type of this rating is %T", input)
	
}