package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
	}
	tree = result[0]
	for idx, node := range result {
		if 2*idx+1 < len(result) {
			node.Left = result[2*idx+1]
		}
		if 2*idx+2 < len(result) {
			node.Right = result[2*idx+2]
		}
	}
	return tree
}

func BenchmarkSol(b *testing.B) {
	root := CreateBinaryTree(&[]string{"3", "1", "4", "null", "2"})
	k := 1
	for idx := 0; idx <= b.N; idx++ {
		kthSmallest(root, k)
	}
}
func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [3,1,4,null,2], k = 1",
			args: args{root: CreateBinaryTree(&[]string{"3", "1", "4", "null", "2"}), k: 1},
			want: 1,
		},
		{
			name: "root = [5,3,6,2,4,null,null,1], k = 3",
			args: args{root: CreateBinaryTree(&[]string{"5", "3", "6", "2", "4", "null", "null", "1"}), k: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
