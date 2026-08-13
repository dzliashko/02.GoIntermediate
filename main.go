package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	// split into 4 chunks, goroutine each, sum total
	var wg sync.WaitGroup
	var total int
	var mu sync.Mutex

	numGoroutines := 4
	wg.Add(numGoroutines)

	for i := range numGoroutines {
		start := i * n / numGoroutines
		end := (i + 1) * n / numGoroutines
		if i == numGoroutines-1 {
			end = n
		}
		go func(chunk []int) {
			defer wg.Done()

			subtotal := 0
			for _, val := range chunk {
				subtotal += val
			}

			mu.Lock()
			total += subtotal
			mu.Unlock()
		}(nums[start:end])
	}
	wg.Wait()

	fmt.Println(total) // replace with the real total
}
