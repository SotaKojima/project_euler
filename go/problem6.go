package main

import (
    "fmt"
)

func pow(num int64) (int64) {
    return num * num
}

func main() {
    var i int64
    var ans int64
    var ans_pow_plus int64 = 0
    var ans_plus_pow int64 = 0

    for i = 1; i <= 100; i++ {
        ans_pow_plus += pow(i)
        ans_plus_pow += i
    }

    ans = pow(ans_plus_pow) - ans_pow_plus

    fmt.Println(ans)
}