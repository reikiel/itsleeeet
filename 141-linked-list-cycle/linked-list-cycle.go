/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Set map[*ListNode]struct{}

func hasCycle(head *ListNode) bool {
    m := make(Set)

    for head != nil {
        if _, ok := m[head]; ok {
            return true
        }
        m[head] = struct{}{}
        head = head.Next
    }

    return false
}