package examples

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/elizabevil/ffmpegx/metadatax"
	"testing"
)

var ss, _ = json.Marshal(metadatax.Progress{})

// options contains a list of &-separated options of the form key=val.

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &metadatax.Progress{}
		json.Unmarshal(ss, stu)
	}
}
func BenchmarkUnmarshalWithPool(b *testing.B) {
	var pool = metadatax.NewProgressPool()
	for n := 0; n < b.N; n++ {
		stu := pool.GetProgress()
		json.Unmarshal(ss, stu)
		stu.Reset()
		pool.PutProgress(stu)
	}
}
func BenchmarkSonicUnmarshalWithPool(b *testing.B) {
	var pool = metadatax.NewProgressPool()
	for n := 0; n < b.N; n++ {
		stu := pool.GetProgress()
		sonic.Unmarshal(ss, stu)
		stu.Reset()
		pool.PutProgress(stu)
	}
}

/*
goos: linux
goarch: amd64
cpu: 13th Gen Intel(R) Core(TM) i9-13980HX
BenchmarkUnmarshal-9                      682942              1689 ns/op             440 B/op          5 allocs/op
BenchmarkUnmarshalWithPool-9             2025644               595.1 ns/op           616 B/op          5 allocs/op
BenchmarkSonicUnmarshalWithPool-9        4362739               275.7 ns/op           672 B/op          5 allocs/op
PASS
ok      command-line-arguments  4.474s
*/
