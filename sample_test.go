package sample

import "testing"

func TestAdd(t *testing.T) {
	expected := 3
	result := Add(1, 2)
	if result != expected {
		t.Errorf("result:%v expected:%v", result, expected)
	}
}
