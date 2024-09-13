/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {return true}

    return depth(root) != -1
}

func depth(root *TreeNode) int {
    if root == nil { return 0 }

    left := depth(root.Left)
    right := depth(root.Right)

    if (left == -1 || right == -1 || AbsInt(left-right) > 1) {
        return -1
    }

    
    return MaxInt(left, right) + 1

}

func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func AbsInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}