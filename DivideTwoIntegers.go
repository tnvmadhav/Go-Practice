// W/o use of Inbuilt variables(math.MaxInt32, etc.) or functions like math.Abs() etc.


func divide(dividend int, divisor int) int {
    sign := 1
    // Handle negative cases
    if dividend < 0 && divisor < 0{
        dividend = -dividend
        divisor = -divisor
    }else{
        if divisor < 0{
            divisor = -divisor
            sign = -1
        } else {
            if dividend < 0{
                dividend = -dividend
                sign = -1
            }
        }
    }
    
    count := 0
    for {
        if dividend < divisor{
            if count*sign > int(math.Pow(float64(2), float64(31)))-1{
                return int(math.Pow(float64(2), float64(31)))-1
            } else {
                if count*sign < -1 * int(math.Pow(float64(2), float64(31))){
                    return -1 * int(math.Pow(float64(2), float64(31)))
                } else{
                    return count*sign
                }
            } 
        }
        dividend -= divisor
        count++
    }
}
