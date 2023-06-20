package main

import "fmt"

var y = 10
var a = "Sweet"

var dialogue =`James said,"He is good at golang"`

func main() {
	fmt.Printf("%T\n",y)
	fmt.Printf("%T\n",a)
	fmt.Println(dialogue)
	fmt.Printf("%T\n",dialogue)
}