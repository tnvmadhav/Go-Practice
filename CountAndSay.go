/* DP Solution
Runtime: 64 ms        :/
Memory Usage: 6.5 MB  :/
*/ 

func DoAll(n string) string {
    x := n
    myMap := make(map[int]int)
    var say []string
    var k int
    for i, _ := range x {
        k++
        if (i == len(x)-1) || (x[i] != x[i+1]){
            myMap[int(x[i]) - '0'] = k
            say = append(say, fmt.Sprintf("%d%d", myMap[int(x[i]) - '0'],int(x[i]) - '0'))
            myMap = make(map[int]int)
            k = 0
        }
    }
    return strings.Join(say, "")
}

func countAndSay(n int) string {
    myArray := []string{"1"}
    for {
        if (len(myArray) > 30){
            break
        }
        myArray = append(myArray, DoAll(myArray[len(myArray)-1]))
    }
    return myArray[n-1]
}
