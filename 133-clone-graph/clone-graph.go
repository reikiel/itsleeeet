/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := make(map[int]*Node)
    return cloneNode(node, visited)
}

func cloneNode(node *Node, visited map[int]*Node) *Node {
    if node == nil {
        return nil
    }

    // if already cloned, return the clone
    if clone, exists := visited[node.Val]; exists {
        return clone
    }

    // clone node and initialise neighbors
    // since we know how many from node
    root := &Node{
        Val: node.Val,
        Neighbors: make([]*Node, len(node.Neighbors)),
    }

    // Store in visited
    visited[node.Val] = root

    // Recursively clone all neighbors
    // add to root.Neighbors at the same time
    for i, neigh := range node.Neighbors {
        root.Neighbors[i] = cloneNode(neigh, visited)
    }

    return root
}