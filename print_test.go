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
	tree := asciitree.Tree{
		"root": nil,
	}

	tree.Fprint(os.Stdout, true, "")
}

func Test_TwoLevels(t *testing.T) {
	tree := asciitree.Tree{
		"root": asciitree.Tree{"sibling": nil},
	}

	tree.Fprint(os.Stdout, true, "")
}

func Test_TwoSiblings(t *testing.T) {
	tree := asciitree.Tree{
		"root": asciitree.Tree{"sibling1": nil, "sibling2": nil},
	}

	tree.Fprint(os.Stdout, true, "")
}

func Test_TwoRoots(t *testing.T) {
	tree := asciitree.Tree{
		"root1": asciitree.Tree{"sibling1": nil, "sibling2": asciitree.Tree{"sibling1": nil, "sibling2": nil}},
		"root2": asciitree.Tree{"sibling1": nil, "sibling2": nil},
	}

	tree.Fprint(os.Stdout, false, "")
}
