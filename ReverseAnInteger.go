/*
** Given a 32-bit signed Integer, return the reversed Integer
** Example: Input -> 123,-123, 100000000000000000000000, -1000000000000000000000000
**          Output -> 321,-321, 0(outside 32[2^31-1] bit range), 0(outside 32[-2^31] bit range)
*/
func reverse(x int) int {
    neg := 1
    if x < 0{
        neg = -1
        x = x * -1
    }
    sum := 0
    xstring := strconv.Itoa(x)
    for i,_ := range xstring{
        val, _ := strconv.Atoi(string(xstring[i]))
        sum = sum + (int(math.Pow(10.0, float64(i))) * val)
    }
    Answer := sum * neg
    if (-1 * int(math.Pow(2.0, float64(31)))  <  Answer && Answer < int(math.Pow(2.0, float64(31))-1)){
        return Answer
    }
        return 0
}
