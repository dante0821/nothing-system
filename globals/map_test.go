package globals

import "testing"

func TestMapToJson(t *testing.T) {
	m := make(map[string]interface{}, 0)
	m["a"] = 10
	m["b"] = 12
	c := make([]int, 0, 10)
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3)
	m["d"] = c
	s := MapToJson(m)
	t.Log(s)
	m2 := JsonToMap(s)
	t.Log(m2)
}
