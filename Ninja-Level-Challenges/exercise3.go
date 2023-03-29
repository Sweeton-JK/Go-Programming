package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main(){

	s:=fmt.Sprintln(x,y,z)
	fmt.Println(s)
	fmt.Printf("%T\n",s)
}