package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[0], arr[1], arr[2])

	var arr1 = [...]int {1, 2, 3}
	t.Log(arr1[0], arr1[1], arr1[2])
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int {1, 2, 3}

	for i := 0; i < len(arr); i ++ {
		t.Log(arr[i])
	}

	for idx, value := range arr {
		t.Log(idx, value)
	}

	for _, value := range arr {
		t.Log(value)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int {1, 2, 3, 4, 5}
    // 取前3个元素 [a:b], index的范围[a, b)
    arr3 := arr[:3]
    t.Log(arr3)


}