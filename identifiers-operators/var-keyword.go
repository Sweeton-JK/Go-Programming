package main 

import "fmt"

var y = "This Is Sweet"
// x:= "Does Not Work" 
var x int

func main() {

	fmt.Println("Hello, Let's Start")
	x = 14
	foo()
}

func foo() {
	fmt.Println(y)
	fmt.Println("Age Is ",x)
}