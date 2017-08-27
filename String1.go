package main

import "fmt"
// if else fo len

func unionStringIndex(x,y,z string)(string) {
	if len(x)>=len(y){
		for i,n:=range y {
			z+=string(n)
			m:=x[i]
			z+=string(m)
		}
	}else{
		for i,n:=range x {
			z+=string(n)
			m:=y[i]
			z+=string(m)
		}
	}
	return(z)
}
func main() {
	var x string = "123456"
	var y string = "abcdetqw"
	var z string
	 fmt.Println(unionStringIndex(x,y,z))
}
