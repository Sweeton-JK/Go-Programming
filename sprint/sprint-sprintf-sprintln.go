package main

import "fmt"

var a int = 10

func main(){

	// x:=fmt.Printf("%T\t%b\t%x\t%#x",a,a,a,a)
	y:=fmt.Sprintf("%T\t%b\t%x\t%#x",a,a,a,a)
 
	// fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",y)
}