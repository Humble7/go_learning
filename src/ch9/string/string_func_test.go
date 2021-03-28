package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "1,2,3"
	parts := strings.Split(s, ",")
	for _,v := range parts {
		t.Log(v)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string value:" + s)

	if v, err := strconv.Atoi("10"); err == nil {
		t.Log(v)
	} else {
		t.Log("convert error!")
	}
}
