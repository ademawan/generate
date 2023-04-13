package main

import "testing"

func BenchmarkLooping(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Looping(test)
	}
}

func BenchmarkLoopingWithRange(b *testing.B) {

	for n := 0; n < b.N; n++ {
		LoopingWithRange(test)
	}
}
