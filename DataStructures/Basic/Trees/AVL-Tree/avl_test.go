package main

import (
	"github.com/stretchr/testify/assert"
	tree "trees.example.com/api/v1"

	"testing"
)

func TestAVLTree(t *testing.T) {
	tests := []struct { // test case definition
		name string
		args bool
		got  any
		want any
	}{ // test cases
		{ // test case
			name: "testing height of empty avl tree",
			args: false,
			got:  tree.NewAVL().Height(),
			want: 0,
		},
		{ // test case
			name: "Insert root data in avl",
			args: false,
			got:  tree.NewAVL().InsertAVL("One", 1).Height(),
			want: 1,
		},
		{ // test case
			name: "Insert root data in avl",
			args: false,
			got:  tree.NewAVL().InsertAVL("One", 1).Height(),
			want: 1,
		},
		{ // test case
			name: "Get the inserted value",
			args: false,
			got:  tree.GetNodeValue(tree.NewAVL().InsertAVL(1, "One").AVLGetRoot()),
			want: "One",
		},
		{ // test case
			name: "Get the inserted key",
			args: false,
			got:  tree.GetNodeKey(tree.NewAVL().InsertAVL(1, "One").AVLGetRoot()),
			want: 1,
		},
		{ // test case
			name: "Bulk insert with single data element",
			args: false,
			got:  tree.BulkInsertAVL(tree.NewAVL(), []any{"Two"}, []any{1}).Size(),
			want: 1,
		},
		{ // test case
			name: "Bulk insert with single data element and get root key",
			args: false,
			got: tree.GetNodeKey(tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two"}, []any{2}).AVLGetRoot()),
			want: 2,
		},
		{ // test case
			name: "Get inOrder data from empty tree",
			args: false,
			got:  tree.NewAVL().GetInOrderValues()[0].(string),
			want: "Empty Tree",
		},
		{ // test case
			name: "Bulk insert into tree and get the root height",
			args: false,
			got: tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "one", "Three", "four", "six", "five"},
				[]any{2, 1, 3, 4, 6, 5}).
				Height(),
			want: 3,
		},
		{ // test case
			name: "Inorder traversal of AVL tre",
			args: false,
			got: tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "one", "Three", "four", "six", "five"},
				[]any{2, 1, 3, 4, 6, 5}).
				GetInOrderValues(),
			want: []any{"one", "Two", "Three", "four", "five", "six"},
		},
	}

	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) { // function call
			assert.Equal(t, tt.want, tt.got) // test case assertions
		})
	}
}
