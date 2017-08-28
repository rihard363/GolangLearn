package main

import( 
  "fmt"
  "strings")
func presenceString(x string)([]string) {
	var sl1 []string
	sl1 = strings.Split(x,"")//преобразование строки в слайс
	return sl1
}
func main() {
  var x string = "12345asd678"
	 fmt.Println(presenceString(x))
}
