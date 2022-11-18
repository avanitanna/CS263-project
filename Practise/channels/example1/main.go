package main

import "fmt"

func main() {

    receive_double_value := make(chan int)
    go func(x int) { receive_double_value <- 2*x }(7) 
    
    val := <-receive_double_value
    fmt.Println(val)

}
