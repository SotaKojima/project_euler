package main

import (
    "fmt"
)

func is_match_num(num int64) (bool) {
    var i int64
    var cnt int64 = 0

    for i = 11; i <= 20; i++ {

        if num % i == 0 {
            cnt++
        }

    }

    if cnt == 10 {
        return true
    }

    return false
}

func main() {
    var i int64

    for i = 1; ; i++ {

        if is_match_num(i) == true {
            break
        }

    }

    fmt.Println(i)
}