package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
	numberOfRunner := 10
	ch := make(chan string, numberOfRunner)
	for i := 0; i < numberOfRunner; i ++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <- ch
}

func AllResponse() string {
	numberOfRunner := 10
	ch := make(chan string, numberOfRunner)
	for i := 0; i < numberOfRunner; i ++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalRet := ""
	for i := 0; i < numberOfRunner; i ++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	//t.Log(FirstResponse())
	t.Log(AllResponse())
	time.Sleep(1 * time.Second)
	t.Log("after:", runtime.NumGoroutine())
}
