package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) speak() {
	fmt.Println("...")
}

func (p *Pet) speakTo(name string) {
	p.speak()
	fmt.Println(name)
}


type Dog struct {
    Pet
}

// 并不能重载
func (d *Dog) speak() {
	fmt.Println("Wang! ")
}

func TestExtension(t *testing.T) {
	d := new(Dog)
	d.speakTo("Mike")
}

