package main

import "testing"

func BenchmarkHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		creatHashTables()
	}
}
