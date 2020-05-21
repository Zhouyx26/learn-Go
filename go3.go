package main

import (
	"fmt"
)
//func main(){
//	fmt.Println("hello")
//	fmt.Println("1+1 = ",1+1)
//	var a string = "1243"
//	fmt.Println(a)
//	// str use  " " two "
// 	b := "4141"
//	fmt.Println("b is a str",b)
//	c,d := 1243,"123"
//
////}
//
//func main(){
//	var i int = 12
//	fmt.Println(i)
//}

func main(){
	s := "hello"
	// 要使用这个才有
	bytes := [5]byte{'1','2','3','4','5'}
	prime := [3]int{2,3,4}
	fmt.Println(s,bytes,prime)

}
const(
	i = 100
	pi = 3.14

)

var frenchHello string
var emptystring  = ""
func test(){
	no,yes,maybe := "no","yes","hello"
	japaneseHello := "Konichiwa"
	frenchHello := "Bonjour"
	fmt.Println(frenchHello,japaneseHello,no,yes,maybe)
}

var s string = "hello"
s[0] = 'c'