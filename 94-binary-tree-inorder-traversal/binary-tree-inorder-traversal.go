/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    res := []int{}

    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        // else
        // pop last off stack and add to res, check if got right
        // if no right then continue loop
        // if have right then add that to the stack, and continue
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, root.Val)
        
        // dont need append here, will be appended in next iteration if not nil
        root = root.Right
    }
    return res
}