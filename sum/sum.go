package main
  
import (
        "fmt"
)

func myWrapper(a []int, sum int, chan1 chan int) {

        chan1<-Sum(a, sum)
        close(chan1)

}

func Sum(a []int, sum int) int {

        if len(a) == 1 {
                return a[0]
                }

        return Sum(a[:len(a)/2], sum) + Sum(a[len(a)/2:], sum)
}


func main() {

        chan1 := make(chan int)
        a := []int{1,2,3,4}
        sum := 0

        go func() {
        fmt.Println(<-chan1)
        }()

        myWrapper(a,sum,chan1)

}
