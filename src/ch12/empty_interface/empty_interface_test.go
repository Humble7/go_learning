package empty_interface

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	if i, success := p.(int); success {
		fmt.Println("Integer Type.", i)
		return
	}
	if s, success := p.(string); success {
		fmt.Println("string Type.", s)
		return
	}

	fmt.Println("Unknown Type.")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	doSomething(10)
	doSomething("10")
	doSomething(false)
}
