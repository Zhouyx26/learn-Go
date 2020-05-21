package main

import "fmt"

//func main(){
//	i := 1
//Here:
//	fmt.Println(i)
//	i++
//	goto Here
//}

//func main(){
//		sum := 0
//		for index:=0;index < 10; index++{
//				sum += index
//		}
//		fmt.Println("sum is equal to",sum)
//
//}

//func main(){
//		for index := 10 ; index > 0;index--{
//				if index==5{
//						break
//				}
//				fmt.Println(index)
//		}
//}

//func main(){
//		for _ , v = range map{
//				fmt.Println("map's key",k)
//				fmt.Println("map's val",v)
//		}
//}


//func main(){
//		integer := 6
//		switch integer {
//		case 4:
//			fmt.Println("the integer was <=4")
//			fallthrough
//		case 6:
//			fmt.Println("the integer ")
//		}
//}

//func max(a,b int) int {
//		if a>b{
//				return a
//		}
//		return b
//}
//
//func main(){
//		x := 3
//		y := 4
//		z := 5
//
//		max_xy := max(x,y)
//		max_xz := max(x,z)
//		fmt.Printf("max(%d,%d) = %d\n",x,y,max_xy)
//		fmt.Printf("max(%d,%d) = %d\n",x,z,max_xz)
//		fmt.Printf("max(%d,%d) = %d\n",y,z,max(y,z))
//
//
//}

//func main(){
//		r := [...]int{99:-1}
//		fmt.Println(r)
//}
//
//func add1(a int) int {
//		a = a+1
//		return a
//}
//
//func main(){
//		x := 3
//
//		fmt.Println("x = ",x)
//		x1 := add1(x)
//		fmt.Println("x+1 = ",x1)
//
//}

func add1(a *int) int {
	*a = *a + 1
	return *a
}

func main(){
	x := 3
	fmt.Println("x=",x)

	x1 := add1(&x)
	fmt.Println("x+1 = ",x1)
	fmt.Println("x = ",x)
}

func SumAndProduct(A,B int)(int,int){
		return A+B , A*B
}

func main(){
		x:=3
		y:=4

		xPLUSy,xTIMESy:= SumAndProduct(x,y)
		fmt.Printf("%d + %d = %d \n",x,y,xPLUSy)
		fmt.Printf("%d * %d = %d \n",x,y,xTIMESy)
}

func ReadWrite()bool{
		file.Open("file")
		if failureX{
				file.Close()
				return false
		}
		if failureY{
				file.Close()
				return false
		}
		file.Close()
		return true
}

func ReadWrite()bool{
		file.Close()
}
for i := 0 ;i<5;i++{
		defer fmt.Printf("%d",i)
}

