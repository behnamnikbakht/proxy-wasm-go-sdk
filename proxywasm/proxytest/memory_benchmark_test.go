//go:build bench
// +build bench

package proxytest

import (
	"testing"
)

/*
go test -bench=BenchmarkMemoryAllocation -benchtime=10000x -benchmem -tags bench  -memprofile memprofile.out
go tool pprof memprofile.out
list CallOnRequestBody
	Output ==>
		Before update:
			2MB        2MB    417:   cs.requestBody = append(cs.requestBodyBuffer, body...)
		After update:
			.   512.01kB    418:   cs.requestBody = appendByCopy(cs.requestBodyBuffer, body)
*/

func BenchmarkMemoryAllocation(b *testing.B) {
	host, reset := NewHostEmulator(NewEmulatorOption().WithVMContext(&testPlugin{buffered: false}))
	defer reset()

	body := []byte("{\"key1\": \"value1\", \"key2\": \"value2\", \"key3\": \"value3\", \"key4\": \"value4\", \"key5\": \"value5\", \"key6\": \"value6\", \"key7\": \"value7\", \"key8\": \"value8\", \"key9\": \"value9\", \"key10\": \"value10\"}")

	for n := 0; n < b.N; n++ {
		id := host.InitializeHttpContext()
		host.CallOnRequestBody(id, body, false)
	}
}
