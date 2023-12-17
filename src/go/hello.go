package main

import "fmt"

func main() {

	questionArry := []string{"what is the smaller", "what is the biger", "What is IOT"}

	for index, value := range questionArry {

		var input string

		fmt.Printf("Index: %d, %s\n", index+1, value)

		fmt.Println("1 cat")
		fmt.Println("2 World")
		fmt.Println("3 the Sea")
		fmt.Println("4 atom")
		fmt.Println("........................................")
		fmt.Println("youre answer is : ")
		fmt.Scanln(&input)

	} 
}
