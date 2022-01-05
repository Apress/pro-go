package main

import (
    "sync"
    "time"    
    // "math"
    // "math/rand"    
    "context"
)

const (
    countKey  = iota
    sleepPeriodKey
)

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
    total := 0
    count := ctx.Value(countKey).(int)
    sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
    for i := 0; i < count; i++ {        
        select {
            case <- ctx.Done():
                if (ctx.Err() == context.Canceled) {
                    Printfln("Stopping processing - request cancelled")
                } else {
                    Printfln("Stopping processing - deadline reached")
                }
                goto end
            default:
                Printfln("Processing request: %v", total)
                total++
                time.Sleep(sleepPeriod)
        }
    }
    Printfln("Request processed...%v", total)
    end: 
    wg.Done()
}

func main() {
    waitGroup := sync.WaitGroup {}
    waitGroup.Add(1)
    Printfln("Request dispatched...")
    ctx, _ := context.WithTimeout(context.Background(), time.Second * 2)
    ctx = context.WithValue(ctx, countKey, 4)
    ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond * 250)
    go processRequest(ctx, &waitGroup)    

    // time.Sleep(time.Second)
    // Printfln("Canceling request")
    // cancel()


    waitGroup.Wait()
}
