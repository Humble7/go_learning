package cancel_by_context

import (
	"context"
	"fmt"
"testing"
"time"
)

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i ++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancel(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

func isCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
