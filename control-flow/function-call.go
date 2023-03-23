package main

import "fmt"

func main() {
	fmt.Println("Hello, I am in Main")
	foo()
	fmt.Println("Back To Main")

	for i :=0 ; i<100 ; i++ {
		
		if i%2==0 {
			fmt.Println("Even Loop")
		}
	}
}

func foo() {
	fmt.Println("Hello, I am in foo")
}