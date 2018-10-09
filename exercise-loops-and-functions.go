package main

import (
        "fmt"
        "math"
)

func Sqrt(x float64) float64 {
        z := float64(x) / 2
        for ; math.Abs(z * z - x) > 0.0005; {
          z -= (z * z - x) / (2 * z)
          fmt.Println("z:", z)
        }
        return z
}

func main() {
        fmt.Println(Sqrt(8))
}
