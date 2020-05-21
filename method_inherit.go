package main

import "fmt"
type Human struct {
		name string
		age int
		phone string

}
type Student struct {
		Human
		School string

}
type Employee struct {
		Human
		company string
}
func (h *Human) SayHi(){
		fmt.Printf("Hi, I'm %s ,you can call me on %s",h.name,h.phone)
}
func (e *Employee) SayHi(){
		fmt.Printf("Hi,I am %s, I work at %s,call me on %s",e.name,e.company,e.phone)
}

func main(){
		mark := Student{Human{"Mark",16,"222-222-YYY"},"MIT"}
		sam := Employee{Human{"Sam",45,"111-999-XXX"},"golang inc"}

		mark.SayHi()
		sam.SayHi()


}