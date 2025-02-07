```go
package main

import (

        "fmt"
        "sync"
        "time"
)

func main() {
        var m sync.Mutex
        wg := sync.WaitGroup{}
        for i := 0; i < 100; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        fmt.Println(i)
                        m.Unlock()
                }(i)
        }
        wg.Wait()
}
```