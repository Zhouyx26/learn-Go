// interface 函数参数
package main
import (
	"fmt"
	"strconv"
)

type Human struct {
		name string
		age int
		phone string
}
type Student struct {
		Human
		school string
		money int
}

func (h Human)String() string{
		return "<"+ h.name +" - "+strconv.Itoa(h.age)+"years - phone"+h.phone+">"
}

func (s Student) String() string{
		return "hello world" + s.name
}
// strconv.Itoa 代表int 转 str
func main(){
		Bob := Student{Human{"Bob",14,"111-222-333"},"baby school",30}
		fmt.Println("This Human is :",Bob)
}

