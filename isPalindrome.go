func isPalindrome(s string) bool {
    var m []string
    l := strings.ToLower(s)
    for _,r := range l{
        if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9'){
            continue
        }
        m = append(m, string(r))
    }
    for i,_ := range m{
        if m[i] != m[len(m)-1-i]{
            return false
        }
    }
    return true
}
