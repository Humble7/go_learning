package _map

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int {1:1, 2:2, 3:3}
	t.Logf("m1 length:%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 4
	t.Logf("m2 length:%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("m3 length:%d", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int {}
	// key不存在会返回0
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	if v,ok := m1[1]; ok {
		t.Log("Key 1's value is", v)
	} else {
		t.Log("Key 1's value is not exists")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int {1:1, 2:2, 3:3}
	for k,v := range m {
		t.Log(k, v)
	}
}

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int {}
	m[0] = func(op int) int {
		return op
	}
	m[1] = func(op int) int {
		return op * op
	}
	m[2] = func(op int) int {
		return op * op * op
	}

	t.Log(m[0](2), m[1](2), m[2](2))
}
