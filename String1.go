package main

import "fmt"

func unionStringByIndex(x, y string)(string) {
	var a string
	if len(x) >= len(y) {
		a = calc(x, y)
	}else {
		a = calc(y, x)
	}
	return(a)
}
func calc(long, short string)(string) {
	var w string
	for i, n := range long {
      if i < len(short) {
        m := short[i]
        w += string(m)
      }
      w += string(n)
	}
	return(w)
}
func main() {
	var x string = "123456789"
	var y string = "abcdetqw"
	 fmt.Println(unionStringByIndex(x, y))
}
// в задаче необходимо сложить строки посимвольно