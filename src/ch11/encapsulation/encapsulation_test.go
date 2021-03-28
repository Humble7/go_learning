package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
    Age int
}

// 为结构添加行为

// 会发生内存拷贝
func (e Employee) String() string {
	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 不会发生内存拷贝
//func (e *Employee) String() string {
//	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 20}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)

	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)

}

func TestMemoryCopy(t *testing.T) {
	e := Employee{"1", "Bob", 20}

	// 观察两次的address是否一样
	t.Log(e.String())
	t.Log(e.String())
}
