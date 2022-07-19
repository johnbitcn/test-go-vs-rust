package main

import "testing"

func TestLoops(t *testing.T) {
	Loops()
}

func BenchmarkLoops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loops()
	}
}
