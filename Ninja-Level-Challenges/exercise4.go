package main

import "fmt"

type hotdog int
var x hotdog = 14

func main(){

	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)

}