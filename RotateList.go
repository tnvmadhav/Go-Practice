/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    var count = 0
    // Null Check
    if head == nil{
        return head
    }
    
    //1.
    t := head
    for {
        count++
        if t.Next == nil{
            break
        }
        t = t.Next
    }
    
    //2.
    t.Next = head
    
    //3.
    l := t
    if k > count{
        k = k % count
    }
    for i := 0; i < count-k; i++{
        l = l.Next
    }
    
    //4.
    head = l.Next
    
    //5.
    l.Next = nil
    
    return head 
}


//1. Get length of linked list by counting while traversing till tail
//2. convert to circular linked list.
//3. Traverse to the new Tail(length - k)th Node
//4. Mark the next node as head
//5. Set the next of tail to Null
