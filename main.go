package main

import (
	"sort"
	"strings"
	"unsafe"
)

func main() {}

//export csv2json
func csvToJson(csv string) uint32 {
	rows := Parse(csv)
	jsonStr := ToJson(rows)

	return stringToPtr(jsonStr)
}

type Row map[string]string
type Rows []Row

func Parse(csv string) (result Rows) {
	line := strings.Split(csv, "\n")
	if len(line) <= 1 {
		print(result)
		return
	}

	header := strings.Split(line[0], ",")
	colsCount := len(header)

	for i := 1; i < len(line); i++ {
		row := strings.Split(line[i], ",")
		if len(row) != colsCount {
			// TODO: ヘッダー行の要素の個数と今見ている行の要素の個数が合わない場合エラーを投げる
		}
		parsed := make(Row, colsCount)
		for j, elem := range row {
			parsed[header[j]] = elem
		}
		result = append(result, parsed)
	}
	return
}

func ToJson(rows Rows) (result string) {
	result = "["
	for i, row := range rows {
		var keys []string
		for key := range row {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		object := "{"
		for j, key := range keys {
			object = object + "\"" + key + "\": " + "\"" + row[key] + "\""
			if j == len(keys)-1 {
				object = object + "}"
			} else {
				object = object + ", "
			}
		}

		result = result + object
		if i == len(rows)-1 {
			result = result + "]"
		} else {
			result = result + ", "
		}
	}
	return
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
