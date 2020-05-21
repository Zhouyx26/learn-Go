package main

import "fmt"

type person struct {
		name string
		age int
}
func main(){
	var P person

	P.name = "alise"
	P.age = 12
	fmt.Printf("the person's name is %s",P.name)

}
