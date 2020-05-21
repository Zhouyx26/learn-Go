package main
import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 15, 55}, "Computer Science"}
	fmt.Println("his name is ", mark.name)
	fmt.Println("his age is ", mark.age)
	fmt.Println("his weight is ", mark.weight)
	fmt.Println("his speciality is ", mark.speciality)

	mark.speciality = "AI"
	fmt.Println("Mark change his speciality")
	fmt.Println("His speciality is ",mark.speciality)
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is ",mark.age)
	// change weight
	fmt.Println("Mark is not athlet anymore ")
	mark.weight += 60
	fmt.Println("His weight is ",mark.weight)

}

