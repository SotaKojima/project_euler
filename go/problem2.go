package main

import (
    "fmt"
)

func fibonacci(num int32) (int32) {
    var pre_num int32 = 0
    var cur_num int32 = 1
    var next_num int32 = 0
    var i int32 = 0

    for i = 0; i < num; i++ {
        next_num = pre_num + cur_num
        pre_num = cur_num
        cur_num = next_num
    }

    return next_num
}

func main() {
    var sum int32 = 0
    var i int32 = 0

    for i = 0; fibonacci(i) <= 4000000; i++ {

        if fibonacci(i) % 2 == 0 {
            sum += fibonacci(i)
        }

    }

    fmt.Println(sum)
}