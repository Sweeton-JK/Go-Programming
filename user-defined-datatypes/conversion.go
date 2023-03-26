package main

import "fmt"

var a int = 10
type hotdog int
var b hotdog = 15

func main(){
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	a=int(b)
	fmt.Println(a)
}