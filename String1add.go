package main

import "fmt"

func unionStringByIndex(x, y string)(string) {
	var w string
	if len(x)>len(y){
		for i :=0 ; i < len(x) ; i ++ {
			n := x[i]
			w += string(n)
			if i < len(y){
				m := y[i]
				w += string(m)
			}
		}
	} else {
		for i :=0 ; i < len(y) ; i ++ {
			if i < len(x){
				n:= x[i]
				w += string(n)
			}
			m := y[i]
			w += string(m)
		}
	}
	return(w)
}
func main() {
	var x string = "123456789"
	var y string = "abcdetqwzxcvb"
	 fmt.Println(unionStringByIndex(x, y))
}
// в задаче необходимо сложить строки посимвольно
// идет индексация по первой строке 
//если числа длинее букв то проход по числам добавляя значение числа потом буквы, если буквы длинее то проход по буквам добавляя сначала число потом букву 