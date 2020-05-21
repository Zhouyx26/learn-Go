package main

import (
	"time"
)

//ch := make(chan type,value)

// buffered channels
//import "fmt"
//func main(){
//		c := make(chan int,2)
//		c <- 4
//		c <- 2
//		c <- 3
//		fmt.Println(<-c)
//		fmt.Println(<-c)
//		fmt.Println(<-c)
//}

// 斐波那契函数
//func fabonacci(n int,c chan int){
//		x,y := 1,1
//		for i := 0; i<n ; i++{
//				c <- x
//				x,y = y,x+y
//		}
//		close(c)
//}
//
//func main(){
//		c := make(chan int,10)
//		go fabonacci(cap(c),c)
//		for i := range c{
//				fmt.Println(i)
//		}
//}

// select

//import "fmt"
//func fibonacci(c ,quit chan int){
//		x,y := 1,1
//		for {
//				select {
//				case c<- x:
//						x,y = y , x+y
//				case <- quit:
//						fmt.Println("quit")
//						return
//				}
//		}
//}
//
//func main(){
//		c := make(chan int)
//		quit := make(chan int)
//		go func(){
//				for i := 0;i<10;i++{
//						fmt.Println(<-c)
//
//				}
//				quit <- 0
//		}()
//		fibonacci(c,quit)
//}
//
//select{
//case i := <-c:
//	//use i
//default:
//	// 当 c 阻塞的时候 执行这里
//}

func main(){
		c := make(chan int)
		o := make(chan bool)
		go func(){
				for {
					select {
							case v := <-c:
								println(v)
							case <- time.After(5*time.Second):
									println("timeout")
									o <- true
									break
						}
				}
		}()
		<- o
}

//
//func main(){
//	go func(){
//		for{
//			select {
//			case  :
//				println("")
//			case :
//				println("")
//
//			}
//
//		}
//	}()
//	<- o
//}
