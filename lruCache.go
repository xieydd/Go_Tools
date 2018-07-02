//LRU算法，即最近最少使用，用于对于不经常使用的数据进行清除

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/Go_Tools/hashtable"
)

type Fib func(int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timeAndCacheFib(f Fib) Fib {
	return func(n int64) int64 {

		defer func(t time.Time) {
			fmt.Printf("---EveryTime Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())

		h := hashtable.NewHash(100)
		result := h.Get(n)
		if result == nil {
			result := f(n)
			h.Add(n, result)
		}

		return f(n)
	}
}

func Fib1(i int64) int64 {
	if i < 2 {
		return 1
	}
	return Fib1(i-1) + Fib1(i-2)
}

func main() {
	var i int64
	t1 := time.Now()
	for i = 0; i < 1000; i++ {
		fib := timeAndCacheFib(Fib1)
		fmt.Printf("%d\n", fib(i))
	}
	elapsed := time.Since(t1)
	fmt.Printf("AllTime Elapsed : %v \n", elapsed)
}
