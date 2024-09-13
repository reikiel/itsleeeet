/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // if one is < root and other > root, or if either == root
    // then thats the LCA
    if root.Val == p.Val || root.Val == q.Val || (q.Val < root.Val && root.Val < p.Val) || (p.Val < root.Val && root.Val < q.Val) {
        return root
    } else if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else { // both is more
        return lowestCommonAncestor(root.Right, p, q)
    }
	
}