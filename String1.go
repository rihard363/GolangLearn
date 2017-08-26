package main

import "fmt"

func unionStringIndex() {
	x:= "123"
	y:= "abc"
	var z string
	for i,n:=range y {
		z+=string(n)
		m:=x[i]
		z+=string(m)
	}
	fmt.Println(z)
}
func main() {
	 unionStringIndex()
}
