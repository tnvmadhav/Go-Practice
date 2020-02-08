// Runtime: 128 ms
// Memory:  7.5 MB

func sortedSquares(A []int) []int {
    myMap := make(map[int]int)
    var keys []int
    var outputArray []int
    
    for i, _ := range A{
        myMap[int(math.Pow(float64(A[i]), 2.0))]++
    }
    
    for k := range myMap{
        keys = append(keys, k)
    }
    
    sort.Ints(keys)
    
    for i := range keys{
        for j:=1; j<=myMap[keys[i]]; j++{
            outputArray = append(outputArray, keys[i])
        }
    }
    return outputArray
}
