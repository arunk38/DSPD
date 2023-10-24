package main

import (
	"github.com/stretchr/testify/assert"
	tree "trees.example.com/api/v1"

	"testing"
)

func TestAVLTree(t *testing.T) {
	tests := []struct { // test case definition
		name string
		skip bool
		got  any
		want any
	}{ // test cases
		{ // test case
			name: "testing height of empty avl tree",
			skip: false,
			got:  tree.NewAVL().Height(),
			want: 0,
		},
		{ // test case
			name: "Insert root data in avl",
			skip: false,
			got:  tree.NewAVL().InsertAVL("One", 1).Height(),
			want: 1,
		},
		{ // test case
			name: "Insert root data in avl",
			skip: false,
			got:  tree.NewAVL().InsertAVL("One", 1).Height(),
			want: 1,
		},
		{ // test case
			name: "Get the inserted value",
			skip: false,
			got:  tree.GetNodeValue(tree.NewAVL().InsertAVL(1, "One").AVLGetRoot()),
			want: "One",
		},
		{ // test case
			name: "Get the inserted key",
			skip: false,
			got:  tree.GetNodeKey(tree.NewAVL().InsertAVL(1, "One").AVLGetRoot()),
			want: 1,
		},
		{ // test case
			name: "Bulk insert with single data element",
			skip: false,
			got:  tree.BulkInsertAVL(tree.NewAVL(), []any{"Two"}, []any{1}).Size(),
			want: 1,
		},
		{ // test case
			name: "Bulk insert with single data element and get root key",
			skip: false,
			got: tree.GetNodeKey(tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two"}, []any{2}).AVLGetRoot()),
			want: 2,
		},
		{ // test case
			name: "Get inOrder data from empty tree",
			skip: false,
			got:  tree.NewAVL().GetInOrderValues()[0].(string),
			want: "Empty Tree",
		},
		{ // test case
			name: "Bulk insert into tree and get the root height",
			skip: false,
			got: tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "one", "Three", "four", "six", "five"},
				[]any{2, 1, 3, 4, 6, 5}).
				Height(),
			want: 3,
		},
		{ // test case
			name: "Inorder traversal of AVL tre",
			skip: false,
			got: tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "one", "Three", "four", "six", "five"},
				[]any{2, 1, 3, 4, 6, 5}).
				GetInOrderValues(),
			want: []any{"one", "Two", "Three", "four", "five", "six"},
		},
		{ // test case
			name: "test root as 4",
			skip: false,
			got: tree.GetNodeValue(tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "one", "Three", "four", "six", "five"},
				[]any{2, 1, 3, 4, 6, 5}).
				AVLGetRoot()),
			want: "four",
		},
		{ // test case
			name: "Bulk insert into tree and do valid search",
			skip: false,
			got: tree.GetNodeValue(
				tree.BulkInsertAVL(tree.NewAVL(), []any{"Two", "one", "Three", "four", "six", "five"}, []any{2, 1, 3, 4, 6, 5}).
					Search(5)),
			want: "five",
		},
		{ // test case
			name: "Get the hieght of any node to validate",
			skip: false,
			got: tree.BulkInsertAVL(tree.NewAVL(), []any{"Two", "one", "Three", "four", "six", "five"}, []any{2, 1, 3, 4, 6, 5}).
				Search(6).GetAVLNodeHeight(),
			want: 2,
		},
		{ // test case
			name: "Bulk insert into tree and do invalid search",
			skip: false,
			got: tree.GetNodeValue(
				tree.BulkInsertAVL(tree.NewAVL(), []any{"Two", "one", "Three", "four", "six", "five"}, []any{2, 1, 3, 4, 6, 5}).
					Search(15)),
			want: "<nil>",
		},
		{ // test case
			name: "Bulk insert into tree and Delete a valid key",
			skip: false,
			got: tree.GetNodeValue(
				tree.BulkInsertAVL(tree.NewAVL(), []any{"Two", "one", "Three", "four", "six", "five"}, []any{2, 1, 3, 4, 6, 5}).
					Delete(1).
					Search(1)),
			want: "<nil>",
		},
		{ // test case
			name: "Bulk insert into tree and Delete a valid key, do in order traversal",
			skip: false,
			got: tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "one", "Three", "four", "six", "five"},
				[]any{2, 1, 3, 4, 6, 5}).
				Delete(1).Delete(6).Delete(4).PrintAVL(),
			want: 0,
		},
		{ // test case
			name: "Print the AVL tress",
			skip: true,
			got: tree.BulkInsertAVL(tree.NewAVL(),
				[]any{"Two", "One", "Three", "Four", "Six", "Five"},
				[]any{2, 1, 3, 4, 6, 5}).PrintAVL(),
			want: 0,
		},
	}

	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) { // function call
			if tt.skip == false {
				assert.Equal(t, tt.want, tt.got) // test case assertions
			}
		})
	}
}
