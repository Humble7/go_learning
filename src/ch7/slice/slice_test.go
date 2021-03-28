package slice

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

    s1 := []int {1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

    s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
    // out of range
    //t.Log(s2[3])
    s2 = append(s2, 1)
    t.Log(s2[3])
}

func TestSliceGrowign(t *testing.T) {
	var slice = []int {}
	for i := 0; i < 10; i ++ {
		slice = append(slice, i)
		t.Log(len(slice), cap(slice))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(len(Q2), cap(Q2))

	Summer := year[5:8]
	t.Log(len(Summer), cap(Summer))
	// 修改了共享数据
    Summer[0] = "unknown"

    t.Log(Q2)
    t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	s1 := []int {1,2,3}
	s2 := []int {1,2,3}

	// slice 不能做比较
	//s1 == s2
}
