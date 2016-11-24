// Concat_test
package main

import (
	"testing"
)

func BenchmarkConcatSprintf(b *testing.B) {
	b.ReportAllocs()
	rec := Record{A: "a", B: "b", C: "c", D: "d", E: "e", F: "f"}
	for i := 0; i < b.N; i++ {
		ConcatSprintf(rec)
	}
}
func BenchmarkConcatJoin(b *testing.B) {
	b.ReportAllocs()
	rec := Record{A: "a", B: "b", C: "c", D: "d", E: "e", F: "f"}
	for i := 0; i < b.N; i++ {
		ConcatJoin(rec)
	}
}
func BenchmarkConcatBuffer(b *testing.B) {
	b.ReportAllocs()
	rec := Record{A: "a", B: "b", C: "c", D: "d", E: "e", F: "f"}
	for i := 0; i < b.N; i++ {
		ConcatBuffer(rec)
	}
}

/*
BenchmarkConcatSprintf-8   	 2000000	       642 ns/op	     112 B/op	       7 allocs/op
BenchmarkConcatJoin-8      	10000000	       133 ns/op	      32 B/op	       2 allocs/op
BenchmarkConcatBuffer-8    	10000000	       188 ns/op	     120 B/op	       2 allocs/op
*/
