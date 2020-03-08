/**
 * Runtime: 4 ms
 * Memory Usage: 2.2 MB
 */
 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil{
        return true
    }
    if q == nil{
        return false
    }
    if p == nil{
        return false
    }
        
    var myArray1,myArray2 = []int{}, []int{}
    myArray1 = popu(p, myArray1)
    myArray2 = popu(q, myArray2)
    fmt.Println(myArray1,myArray2)
    
    if len(myArray1) != len(myArray2){
            return false
    } else {
        for i := range myArray1{
            if myArray1[i] != myArray2[i]{
                return false
            }
        }
    }
    return true
}


func popu(p *TreeNode, a []int) []int {
    a = append(a, p.Val)
    if p.Left != nil{
         a = popu(p.Left, a)  
    } else {
        a = append(a, -1)
    }
    if p.Right != nil{
       a = popu(p.Right, a) 
    } else {
        a = append(a, -1)
    }
    return a
}
