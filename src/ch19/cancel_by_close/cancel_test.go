package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	cancelCh := make(chan struct{}, 0)
	for i := 0; i < 5; i ++ {
		go func(i int, cancelCh chan struct{}) {
            for {
            	if isCancel(cancelCh) {
                    break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelCh)
	}
	//cancel(cancelCh)
	cancel_1(cancelCh)
	time.Sleep(time.Second * 1)
}

func cancel(ch chan  struct{}) {
	ch <- struct{}{}
}

func cancel_1(ch chan  struct{}) {
	close(ch)
}

func isCancel(cancel chan struct{}) bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}
