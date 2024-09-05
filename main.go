package main

import "fmt"
import "time"

func main(){
    ch := make(chan int, 2)
    go first(ch)
    //<-ch
    go second(ch)
    //select {
    <- ch 
    //}
    fmt.Println("run")
    time.Sleep(2 * time.Second)
}


func first(ch chan int){
    for i := 0 ; i <  5 ; i ++ {
        fmt.Println(i)
    }
    ch <- 1
}

func second(ch chan int){

    for i := 5 ; i < 10 ; i ++ {
        fmt.Println(i)
    }
   ch  <- 2 
}
