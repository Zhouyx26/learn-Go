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
		loan float32
}

type Employee struct {
		Human
		company string
		money float32
}
// human 实现 sayhi 方法
func (h Human) SayHi(){
		fmt.Printf("hi I'm %s you can call me on %s\n",h.name,h.phone)

}
func (h Human) Sing(lyrics string){
		fmt.Println("La la la la ...",lyrics)
}
func (e Employee) SayHi(){
		fmt.Printf("Hi I am %s ,I work at %s.Call me on %s\n",e.name,e.company,e.phone)
}

type Men interface {
		SayHi()
		Sing(lyrics string)
}

func main(){
		mike := Student{Human{"Mike",24,"111-111-XXX"},"MIT",0.00}
		paul := Student{Human{"Paul",15,"222-222-XXX"},"Harvard",15}
		sam := Employee{Human{"Sam",37,"333-333-XXX"},"Golang Inc",1500}
		tom := Employee{Human{"Tom",40,"444-444-XXX"},"Thing Ltd",5000}

		var i Men
		// i 储存 student
		i = mike
		fmt.Println("This is Mike,a Student:")
		i.SayHi()
		i.Sing("Moon river ~ ~")

		// i 储存 employee
		i = tom
		fmt.Println("this is Tom,an Employee:")
		i.SayHi()
		i.Sing("Noting left to loss ~ ~ ")

		//定义 slice man
		fmt.Println("Let's use a slice of men and see what happens")
		x := make([]Men,3)
		// 三个都是不同类型的元素，但是他们实现乐interface的同一个接口
		x[0],x[1],x[2] = paul,sam,mike

		for _, value := range x{
				value.SayHi()
		}
}