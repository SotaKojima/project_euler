package main

import (
    "fmt"
    "strconv"
)

func get_num_len(num int64) (int) {
    return len(strconv.FormatInt(num, 10))
}

func is_even_len_num(num int64) (bool) {
    
    if get_num_len(num) % 2 == 0 {
        return true
    }

    return false
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

func reverse_str_num(num int64) (string) {
    var result_str string = ""
    var i int
    var str string = strconv.FormatInt(num, 10)

    for i = get_num_len(num); i > 0; i-- {
        result_str += str[i - 1:i]
    }

    return result_str
}

func add_zero(str_num string, zero_num int64) (string) {
    var i int64

    for i = 0; i < zero_num; i++ {
        str_num += "0"
    }

    return str_num
}

func is_circule_num(num int64) (bool) {
    var right_num string
    var left_num string
    var tmp int64
    var str_len int64
    
    if is_even_len_num(num) == true {
        right_num = strconv.FormatInt(num, 10)[:get_num_len(num) / 2]
        tmp, _ = strconv.ParseInt(strconv.FormatInt(num, 10)[get_num_len(num) / 2:], 10, 64)
        left_num = reverse_str_num(tmp)

        str_len = int64(len(strconv.FormatInt(num, 10)[get_num_len(num) / 2:]))
        left_num = add_zero(left_num, str_len - int64(len(left_num)))

        if right_num == left_num {
            return true
        }
    } else {
        right_num = strconv.FormatInt(num, 10)[:get_num_len(num) / 2 + 1]
        tmp, _ = strconv.ParseInt(strconv.FormatInt(num, 10)[get_num_len(num) / 2 + 1:], 10, 64)
        left_num = reverse_str_num(tmp)

        str_len = int64(len(strconv.FormatInt(num, 10)[get_num_len(num) / 2 + 1:]))
        left_num = add_zero(left_num, str_len - int64(len(left_num)))

        if right_num == left_num {
            return true
        }
    }

    return false
}

func main() {
    var result_list []int64
    var i int64
    var j int64

    for i = 1; i < 1000; i++ {

        for j = 1; j < 1000; j++ {

            if is_circule_num(i * j) == true {
                result_list = append(result_list, i * j)
            }

        }

    }

    fmt.Println(get_max_num(result_list))
}