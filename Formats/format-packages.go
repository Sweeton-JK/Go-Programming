package main

import "fmt"

var y int = 5

func main(){
	fmt.Println(y)
	fmt.Printf("%T\n",y)
	fmt.Printf("%b\n",y)
	fmt.Printf("%x\n",y)
	fmt.Printf("%#x\n",y)
}