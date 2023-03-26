package main

import "fmt"

var a int = 14
type sweet string
var str sweet ="Hello"

func main(){
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",str)
}