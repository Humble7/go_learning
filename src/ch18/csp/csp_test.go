package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("Working on something else.")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	t.Log(service())
	otherTask()
}

func AsyncService() chan string {
	// 方式一
	retCh := make(chan string)

	// 方式二
	//retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret

		// 如果是方式一的Chanel，下面这条语句会等到从Chanel中取出内容后才会执行
		// 方式二会立马执行
		fmt.Println("service existed")
	}()

	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	t.Log(<- retCh)
	time.Sleep(time.Second * 1)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 40):
		t.Error("time out")

	}
}
