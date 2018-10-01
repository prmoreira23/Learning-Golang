package main

import (
    "fmt"
    "math/cmplx"
)

var (
    toBe bool = false
    MaxInt uint64 = 1 << 64 -1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
    fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}
