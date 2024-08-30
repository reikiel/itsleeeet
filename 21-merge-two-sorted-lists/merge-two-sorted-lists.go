/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if (list1 == nil && list2 == nil) { return nil }
    if (list2 == nil && list1 != nil) { return list1 }
    if (list1 == nil && list2 != nil) { return list2 }

    head := new(ListNode)
    move := head
    for (list1 != nil && list2 != nil) {
        if (list1.Val <= list2.Val) {
            move.Next = list1
            move = move.Next
            list1 = list1.Next
        } else {
            move.Next = list2
            move = move.Next
            list2 = list2.Next
        }
    }

    if (list1 != nil) {move.Next = list1}
    if (list2 != nil) {move.Next = list2}

    return head.Next
}