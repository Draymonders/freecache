package main

import (
    "fmt"
    "github.com/coocood/freecache"
    "strings"
    "time"
)

func main() {
    cache := freecache.NewCache(1024 * 50) // 50 KB
    expireTime := 5
    var err error

    for i := 1; i < 5; i++ {
        k, v := kv(i)
        err = cache.Set([]byte(k), []byte(v), 0)
        if err != nil {
            fmt.Printf("key %v set fail, err: %v", k, err)
        }
    }
    expireK, expireV := kv(10)
    err = cache.Set([]byte(expireK), []byte(expireV), expireTime)
    if err != nil {
        fmt.Printf("key %v set expire fail, err: %v", expireK, err)
    }

    k, v := kv(1)
    val, err := cache.Get([]byte(k))
    if err != nil {
        fmt.Printf("key %v get fail, err: %v", k, err)
    }
    if !eq([]byte(v), val) {
        fmt.Printf("key %v get val is not generate val", k)
    }
    time.Sleep(time.Duration(expireTime) * time.Second)
    _, err = cache.Get([]byte(expireK))
    if err != freecache.ErrNotFound {
        fmt.Printf("key %v get fail", expireK, err)
    }
    fmt.Println("look good to me")
}

func kv(i int) (key string, val string){
    key = fmt.Sprintf("key#%d", i)
    val = strings.Repeat(key, 10)
    return
}

func eq(a, b []byte) bool {
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}