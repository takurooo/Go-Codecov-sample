package sample

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		left     int
		right    int
		expected int
	}{
		{name: "1+1", left: 1, right: 1, expected: 2},
		{name: "1+10", left: 1, right: 10, expected: 11},
	}
	for _, test := range tests {
		result := Add(test.left, test.right)
		if result != test.expected {
			t.Errorf("name:%v result:%v expected:%v", test.name, result, test.expected)
		}
	}
}
