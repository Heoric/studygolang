package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var state = make(map[int]int)

    var mutex = &sync.Mutex{}

    var ops int64 = 0

    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                automic.AddInt64(&ops, 1)

                runtime.Gosched()
            }()
        }
    }
}
