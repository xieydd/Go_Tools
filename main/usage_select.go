package main

import (
    "fmt"
    "time"
)

var (
    ch = make(chan string)
    ch1 = make(chan string)
    ch2 = make(chan string)
)

func process(ch chan string) {
    time.Sleep(5*time.Second)
    ch <- "process successful"
}

//Usage of select , when we can receive the distribution response, we can get the most quickest response and ignore the rest. 
func test1() {
    go process(ch)

    for {
        time.Sleep(time.Second)
        select {
        case v := <-ch:
            fmt.Println("Receive the value", v)
            return
        default:
            fmt.Println("No value received")
        }
    }
}


func server1(ch chan string) {
    time.Sleep(time.Second)
    ch <- "Server1 response"
}

func server2(ch chan string) {
    time.Sleep(time.Second)
    ch <- "Server2 response"
}

func test2() {
   go server1(ch1)
   go server2(ch2)

   for {
        select {
        case s1 := <-ch1:
            fmt.Println(s1)
            return
        case s2:= <-ch2:
            fmt.Println(s2)
            return
        default:
            fmt.Println("Nothing")
        }
   }
}

func main() {
    test2()
}
