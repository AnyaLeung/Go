package main

import "fmt"

func vals() (int, int){
    return 3, 7
}

func main(){
    v1, v2 := vals() 
    fmt.Println(v1)
    fmt.Println(v2)

    _, v3 := vals()
    fmt.Println(v3)
}
