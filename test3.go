package main

var ar = [10]byte {'a','b','v','d','r','g','q','w','t'.'b'}
var a,b []byte

a = ar[2:6]
b = ar[3:5]
c = ar[0:len(ar)]

// map 不能用索引取得，只能用键去获取值
var numbers map[string]int
numbers = make(map[string]int)
numbers["one"] = 1
numbers["ten"] = 10
numbers["three"] = 3

rating := map[string]float32{
	"c":5,
	"Go":4.5,
	"Python":4.5,
	"C++":2
}

csharpRating,ok = rating["C#"]

var number map[string]int

number = make(map[string]int)

if x > 10 {
	fmt.Println("this is greater than 10")
}else{
	fmt.Println("this a ")
}

if x := computedValue();x>10{
		fmt.P
}

if x>10{
	fmt.d
}else if {
	fmt.
}else{

}


func myFunc()