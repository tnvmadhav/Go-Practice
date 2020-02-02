//https://leetcode.com/problems/string-to-integer-atoi/ -> :O

func cleanString(str string) string {
    point := 0
    for i, _ := range str{
        if (str[i] == ' '){
            continue
        } else {
            point = i
            break
        }
    }
    return str[point:]
}


func checkValid(str string) (bool, int, string) {
    if len(str) == 0{
        return false, 1, str
    }
    if (str[0] >= '0' && str[0] <= '9'){
        return true, 1, str
    }
    if (str[0] == '-' || str[0] == '+'){
        if (len(str) == 1){
            return false, 1, str
        } else {
            if (str[1] >= '0' && str[1] <= '9'){
                switch str[0]{
                    case '-':
                        return true, -1, str[1:]
                    case '+':
                        return true, 1, str[1:]
                }
            }
        }
    }
    return false, 1, str
}

func CheckLimit(a int) int {
    if (a > int(math.Pow(2.0, float64(31))-1)){
        return int(math.Pow(2.0, float64(31))-1)
    }
    if (a < -1 * int(math.Pow(2.0, float64(31)))){
        return -1 * int(math.Pow(2.0, float64(31)))
    }
    return a
}


func myAtoi(str string) int {
    //1. first remove preceding whitespaces
    cs := cleanString(str)
    //2. Check for first character as a valid numeric digit or '-'
    cv, neg, fine := checkValid(cs)
    fmt.Println(fine)
    if cv {
        k := len(fine)-1
        for i,_ := range string(fine){
            if (fine[i] > '9' || fine[i] < '0'){
                k = i-1
                break
            }
        }
        answer,_ := strconv.Atoi(string(fine[:k+1]))
        //3. Get the number if valid else return 0
        return CheckLimit(answer*neg)
    }
    return 0
}
