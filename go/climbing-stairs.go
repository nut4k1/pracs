// 70. Climbing Stairs
package main

import (
	"fmt"
)

func climbStairs(n int) int {
    var cache = map[int]int { 0:0, 1:1, 2:2, 3:3 }

    return climbStairsWithCache(n, cache)
}

func climbStairsWithCache(n int, cache map[int]int) int {
    if val, ok := cache[n]; ok {
        return val
    }
    cache[n-1] = climbStairsWithCache(n - 1, cache)
    cache[n-2] = climbStairsWithCache(n - 2, cache)
    return cache[n-1] + cache[n-2]
}

func main() {
	fmt.Println(climbStairs(16))
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(45))
}
