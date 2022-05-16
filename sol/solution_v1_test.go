package sol

import "testing"

func BenchmarkSolV1(b *testing.B) {
	root := CreateBinaryTree(&[]string{"3", "1", "4", "null", "2"})
	k := 1
	for idx := 0; idx <= b.N; idx++ {
		kthSmallestV1(root, k)
	}
}
func Test_kthSmallestV1(t *testing.T) {
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
			if got := kthSmallestV1(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
