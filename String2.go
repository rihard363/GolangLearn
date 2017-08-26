package main

import ("fmt"
"strings")

func presenceString() {
	x:= "wertyuiopertyuiopqwerty"
	y:= "qwerty"
	var z string
	z = strings.Replace(x, y, "",-1)
	fmt.Println(z)
}
func main() {
	 presenceString()
}
