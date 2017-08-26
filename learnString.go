package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	slice2 := []string{"a","b","c",}
	var slice3 []string
	for i:=0; i<len(slice1); i++ {
		slice3 = append(slice3,string(slice1[i]+48))
	}
	for n:=0; n<len(slice3); n++ {
		slice3[n] = slice3[n] + slice2[n]
	}
	fmt.Println(slice3)
}
