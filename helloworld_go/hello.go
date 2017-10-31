package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your username? : ")
	user, _ := reader.ReadString('\n')

	if user == "dave\n" {
		fmt.Println("Hello, " + user + "This is your computer.")
	} else {
		fmt.Println("Hello, " + user + "This is not your computer.")
	}

	/*fmt.Println("Let's test a for loop printing 1-10")
	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}
	*/
}
