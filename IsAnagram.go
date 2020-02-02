func isAnagram(s string, t string) bool {
    myMap := make(map[string]int)
    if len(s) != len(t){
        return false
    }
    for i := 0; i < len(s); i++{
        myMap[string(s[i])]++
        myMap[string(t[i])]--
    }
    for _,v := range myMap {
        if v != 0{
            return false
        }
    }
    return true
}
