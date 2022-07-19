package main

import "testing"

func TestLoop_for_empty(t *testing.T) {
	Loop_for_empty()
}

func TestLoop_for_add(t *testing.T) {
	Loop_for_add()
}

func BenchmarkLoop_for_empty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop_for_empty()
	}
}

func BenchmarkLoop_for_add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop_for_add()
	}
}
