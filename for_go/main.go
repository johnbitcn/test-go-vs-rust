package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func Loops() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	//maxloops := 2 ^ 30
	maxloops := int(math.Pow(2, 30))
	perloops := maxloops / cpus
	addloops := maxloops % cpus
	fmt.Printf("Got CPUs: %d\nMax Loops:%d\nPerLoops:%d\nAddLoops:%d\n",
		cpus, maxloops, perloops, addloops)

	thread := func(st int, en int, wg *sync.WaitGroup) {
		for i := st; i < en; i++ {
			sum := 0
			for n := i; n > 0; n >>= 1 {
				if n&1 == 1 {
					sum += 1
				}
			}
			//if sum == 29 {
			//	fmt.Println("Found Number 29!")
			//}
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

func main() {
	fmt.Println("函数运行开始！")
	start := time.Now()
	Loops()
	elapsed := time.Since(start)
	fmt.Printf("函数运行时间：%v", elapsed)
}
