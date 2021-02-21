package main

import//(
    "fmt"
//    "reflect"
//)

func sum(nums ...int){
 //   fmt.Println("type:", reflect.TypeOf(nums))
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums{
        total += num
    }
    fmt.Println(total)
}

func main(){
    sum(1, 2)
    
//    fmt.Println("-------------")

    nums := []int{1, 2, 3}
    sum(nums...)
}
