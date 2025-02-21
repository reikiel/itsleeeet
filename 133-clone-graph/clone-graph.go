/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := [101]bool{}
    queue := []*Node{node} // queue contains old nodes
    cloneMap := make(map[int]*Node) // map contains new nodes

    cloneMap[1] = &Node{ 
        Val: 1,
    }

    for len(queue) > 0 {
        // pop from queue
        currOld := queue[0]
        queue = queue[1:]

        // if visited continue?
        if visited[currOld.Val] {
            continue
        }

        currClone := cloneMap[currOld.Val]

        for _, n := range currOld.Neighbors {
            // if map[n.Val], use that as already cloned. 
            // else make a clone of n and add to map
            var newNeigh *Node
            if _, ok := cloneMap[n.Val]; ok {
                newNeigh = cloneMap[n.Val]
            } else {
                newNeigh = &Node{
                    Val: n.Val,
                }
                cloneMap[n.Val] = newNeigh
            }
            // add to newNode.Neighbours
            // push n to queue
            currClone.Neighbors = append(currClone.Neighbors, newNeigh)
            queue = append(queue, n)
        }
        // mark current node visited
        visited[currOld.Val] = true
        
    }
    
    return cloneMap[1]
}