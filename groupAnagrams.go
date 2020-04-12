import (
    "sort"
)
func groupAnagrams(strs []string) [][]string {
    var myDict = make(map[string][]string)
    for i := range strs {
        myDict[conv(strs[i])] = append( myDict[conv(strs[i])], strs[i])
    }
    
    answer := [][]string{}
    
    for i := range myDict {
        answer = append(answer, myDict[i])
    }
    return answer
}

func conv(s string) string {
    newArr := strings.Split(s, "")
    sort.Strings(newArr)
    return strings.Join(newArr,"")
}
