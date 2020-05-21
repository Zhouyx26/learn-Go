package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List [] Element

type Person struct {
		name string
		age int
}

func (p Person) String() string {
		return "(name:"+p.name+"- age:" + strconv.Itoa(p.age)+"years)"
}

//func main() {
//	list := make(List, 3)
//	list[0] = 1
//	list[1] = "Hello"
//	list[2] = Person{"Dennis", 60}
//
//	for index, element := range list {
//		if value, ok := element.(int); ok {
//			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
//		} else if value, ok := element.(string); ok {
//			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
//		} else if value, ok := element.(Person); ok {
//			fmt.Printf("list[%d] is a person and its value is %s\n", index, value)
//		} else {
//			fmt.Printf("list[%d] is of a differnt type \n", index)
//
//		}
//
//	}
//}

// 采用switch 改写代码

func main(){
		x := make(List,3)
		x[0] = 1
		x[1] = "hello"
		x[2] = Person{"Dennis",70}
		for index,element := range x{
				switch value := element.(type){
					case int:
						fmt.Printf("list[%d] is an int and value is %d\n",index,value)
					case string:
						fmt.Printf("list[%d] is a string and its value is %s\n",index,value)

					case Person:
						fmt.Printf("list[%d] is a Person and its value is %s\n",index,value)
					default:
						fmt.Println("list[%d] is a different type",index)
				}
		}
}