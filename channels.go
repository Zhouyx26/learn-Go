package main
//ci := make(chan int)
//cs := make(chan string)
//cf := make(chan interface{})
//
//ch <- v // 发送 v 到 channel ch.
//v := <-ch //从ch 接受数据
//
import "fmt"
func sum(a []int ,c chan int){
		total := 0
		for _, v := range a {
				total +=v
		}
		c <- total
}
func main(){
		a := []int{7,2,8,-9,4,0}
		c := make(chan int)
		go sum(a[:len(a)/2],c)
		go sum(a[len(a)/2:],c)
		x ,y := <-c,<-c // send total to c
		fmt.Println(x,y,x+y)
}

