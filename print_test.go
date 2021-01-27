package asciitree_test

import (
	"os"
	"testing"

	asciitree "github.com/tufin/asciitree"
)

func Test_EmptyTree(t *testing.T) {
	tree := asciitree.Tree{}

	tree.Fprint(os.Stdout, true, "")
}

func Test_OneLevel(t *testing.T) {
	tree := asciitree.Tree{}
	tree.Add("/root")

	tree.Fprint(os.Stdout, true, "")
}

func Test_TwoLevels(t *testing.T) {
	tree := asciitree.Tree{}
	tree.Add("/root/sibling")

	tree.Fprint(os.Stdout, true, "")
}

func Test_TwoSiblings(t *testing.T) {
	tree := asciitree.Tree{}
	tree.Add("/root/sibling1")
	tree.Add("/root/sibling2")

	tree.Fprint(os.Stdout, true, "")
}

func Test_TwoRoots(t *testing.T) {
	tree := asciitree.Tree{}
	tree.Add("/root/sibling1")
	tree.Add("/root/sibling2/sibling1")
	tree.Add("/root/sibling2/sibling2")
	
	tree.Fprint(os.Stdout, false, "")
}
