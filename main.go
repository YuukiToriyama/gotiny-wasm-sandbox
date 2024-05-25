package main

import (
	"github.com/YuukiToriyama/gotiny-wasm-sandbox/parser"
	"github.com/YuukiToriyama/gotiny-wasm-sandbox/serializer"
	"unsafe"
)

func main() {}

//export csv2json
func csvToJson(csv string) uint32 {
	rows := parser.Parse(csv)
	jsonStr := serializer.ToJson(rows)

	return stringToPtr(jsonStr)
}

// NOTICE: 以下は https://github.com/kajikentaro/wasm-go-tinygo-sample を参考にした。

var wasmBuf [1024]byte

//export getWasmBuffer
func getBuffer() *byte {
	return &wasmBuf[0]
}

var stringBufSize uint32

//export getBufSize
func getBufSize() uint32 {
	return stringBufSize
}

func stringToPtr(s string) uint32 {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))

	stringBufSize = uint32(len(buf))
	return uint32(unsafePtr)
}
