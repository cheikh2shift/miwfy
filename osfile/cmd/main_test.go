package main

import "testing"

func BenchmarkAppendMode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		appendMode()
	}
}

func BenchmarkBruteMode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		bruteMode()
	}
}
