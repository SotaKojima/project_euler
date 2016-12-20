package main

import (
    "fmt"
)

func remove_num(list []int64, num int64) ([]int64) {
    var result_list []int64
    var i int64

    for i = 0; i < int64(len(list)); i++ {
        if list[i] % num != 0 {
            result_list = append(result_list, list[i])
        }
    }

    return result_list
}

func is_include(list []int64, item int64) (bool) {
    var i int64

    for i = 0; i < int64(len(list)); i++ {
        if list[i] == item {
            return true
        }
    }

    return false
}

func sieve_of_eratosthenes(num int64) ([]int64) {
    var prime_num_list []int64
    var num_list []int64
    var i int64

    for i = 2; i <= num; i++ {
        num_list = append(num_list, i)
    }

    for i = 2; i <= num; i++ {

        if is_include(num_list, i) == true {
            prime_num_list = append(prime_num_list, i)
        }

        num_list = remove_num(num_list, i)
    }

    return prime_num_list
}

func get_multiple(list []int64) (int64) {
    var i int64
    var multi int64 = 1

    for i = 0; i < int64(len(list)); i++ {
        multi *= list[i]
    }

    return multi
}

func get_max_num(list []int64) (int64) {
    var i int64
    var max int64 = list[0]

    for i = 0; i < int64(len(list)); i++ {
        
        if max < list[i] {
            max = list[i]
        }

    }

    return max
}

func main() {
    var i int64
    var x int64 = 600851475143
    var div_num int64 = x
    var result_list []int64
    var prime_num_list []int64 = sieve_of_eratosthenes(x)

    for i = 0; i < int64(len(prime_num_list)); i++ {

        if div_num % prime_num_list[i] == 0 {
            result_list = append(result_list, prime_num_list[i])
            div_num /= prime_num_list[i]
            i = 0
        }

        if get_multiple(result_list) == x {
            break
        }

    }

    fmt.Println(get_max_num(result_list))
}