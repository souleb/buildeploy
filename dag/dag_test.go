package dag

import (
	"strings"
	"testing"
)

func TestGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			"empty",
			[]int{1, 2, 3},
			"1\n2\n3\n",
		},
		{
			"test",
			[]int{1, 2, 3},
			"1\n2\n3\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}
			actual := strings.TrimSpace(g.String())
			expected := strings.TrimSpace(tc.expected)
			if actual != expected {
				t.Fatalf("Test %s \n expected: %s\n actual: %s", tc.name, expected, actual)
			}
		})

	}
}
