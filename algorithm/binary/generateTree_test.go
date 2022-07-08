package binary

import "testing"

func TestGenerateTree(t *testing.T) {
	x := generateTrees(3)
	t.Log(x)
}
