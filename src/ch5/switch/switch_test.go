package _switch

import "testing"

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i ++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i ++ {
		switch {
		case i & 1 == 0:
			t.Log("Even")
		case i & 1 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}