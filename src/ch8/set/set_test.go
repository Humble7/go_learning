package set

import "testing"

func TestSetWithMap(t *testing.T) {
	mySet := map[int]bool {}

	if mySet[1] {
		t.Log("1 is existing!")
	} else {
		t.Log("1 is not existing!")
	}

	mySet[1] = true
	if mySet[1] {
		t.Log("1 is existing!")
	} else {
		t.Log("1 is not existing!")
	}

	delete(mySet, 1)
	if mySet[1] {
		t.Log("1 is existing!")
	} else {
		t.Log("1 is not existing!")
	}
}
