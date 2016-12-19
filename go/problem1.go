package main

import (
    "fmt"
)

func is_match_num(num int) (bool) {
    
    if num % 3 == 0 || num % 5 == 0 {
        return true
    }

    return false
}

func main() {
    sum := 0

    for i := 1; i < 1000; i++ {

        if is_match_num(i) == true {
            sum += i
        }

    }

    fmt.Println(sum)
}