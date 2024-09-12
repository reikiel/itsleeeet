/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Filler struct{}

func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]*Filler)

    for head != nil {
        if _, ok := m[head]; ok {
            return true
        }
        m[head] = new(Filler)
        head = head.Next
    }

    return false
}