package sol

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		root TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				root: TreeNode{
					Val: 5,
					L: &TreeNode{
						Val: 3,
						L: &TreeNode{
							Val: -1,
						},
					},
					R: &TreeNode{
						Val: 9,
						L: &TreeNode{
							Val: 6,
							R: &TreeNode{
								Val: 7,
							},
						},
						R: &TreeNode{
							Val: 10,
						},
					},
				},
				k: 4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
