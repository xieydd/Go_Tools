package main

import (
    "sync"
    "fmt"
)

func test1(){
    m := make(map[int]int)
    go func() {
        for {
        _ = m[1]
        }
    }()

    go func() {
        for {
            m[2] = 2
        }
    }()

    // Use standlone will cause Dead Lock
    select {}
}

type mapSync struct {
    sync.RWMutex
    m map[string]int
}

func test2() {
    ma := mapSync{m: make(map[string]int)}
    ma.Lock()
    ma.m["1"]++
    ma.Unlock()

    ma.RLock()
    n := ma.m["1"]
    ma.RUnlock()
    fmt.Println("1:", n)
}

func test3() {
   var m sync.Map

   go func() {
    for i := 0; i<100;i++ {
        m.Store(i,i)
        v,_ := m.Load(i)
        fmt.Printf("%d:\t\t%d\n\t\ti",i,v)
    }
   }()
   go func() {
    for j := 0;j<100;j++ {
        m.Store(j,j)
        s,_ := m.Load(j)
        fmt.Printf("%d:\t\t%d\n\t\tj",j,s)
    }
   }()
   go func() {
    f := func(k,v interface{}) bool {
        fmt.Println("test-Range ",k,v)
        return true
    }
    m.Range(f)
   }()

   select {
   //default:
   // go func() {
   //     k,ok := m.LoadOrStore(1,3)
   //     fmt.Println(k,ok)
   // }()
   }

}

func main() {
    test3()
}
