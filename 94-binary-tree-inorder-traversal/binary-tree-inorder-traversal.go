/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    stack := []*TreeNode{}
    res := []int{}

    stack = append(stack, root)

    for len(stack) > 0 {
        for root.Left != nil {
            stack = append(stack, root.Left)
            temp := root
            root = root.Left
            temp.Left = nil
        }

        // else
        // pop last off stack and add to res, check if got right
        // if no right then continue loop
        // if have right then add that to the stack, and continue
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, root.Val)
        if root.Right != nil {
            stack = append(stack, root.Right)
            root = root.Right
        }

    }
    return res
}