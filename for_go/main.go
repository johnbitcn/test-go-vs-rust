package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func Loops(lens int, tag int) {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	//maxloops := 2 ^ 30
	maxloops := int(math.Pow(2, float64(lens)))
	perloops := maxloops / cpus
	addloops := maxloops % cpus
	fmt.Printf("Got CPUs: %d\nMax Loops:%d\nPerLoops:%d\nAddLoops:%d\n",
		cpus, maxloops, perloops, addloops)

	thread := func(st int, en int, wg *sync.WaitGroup) {
		for i := st; i < en; i++ {
			sum := 0
			nmb := 1
			for n := i; n > 0; n >>= 1 {
				if n&1 == 1 {
					sum += nmb
				}
				nmb++
			}
			if sum == tag {
				fmt.Printf("Got Number is :%d", sum)
				break
			}
		}
		wg.Done()
	}

	var wg sync.WaitGroup

	start := 0
	for cpu := 0; cpu < cpus; cpu++ {
		end := (cpu + 1) * perloops
		if cpu+1 == cpus {
			end += addloops - 1
		}
		fmt.Printf("cpu:%d start:%d end:%d \n", cpu, start, end)
		wg.Add(1)
		go thread(start, end, &wg)
		start += perloops
	}
	wg.Wait()
}

func Empty_loop() int {
	max := int(math.Pow(2, 20))
	v := 0
	for i := 0; i < max; i++ {
		v = i
	}
	return v
}

func main() {
	fmt.Println("函数运行开始！")
	start := time.Now()
	Loops(30, 900)
	//Empty_loop()
	elapsed := time.Since(start)
	fmt.Printf("函数运行时间：%v", elapsed)
}
