package parser

import (
	"github.com/YuukiToriyama/gotiny-wasm-sandbox/types"
	"strings"
)

func Parse(csv string) (result types.Rows) {
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
		parsed := make(types.Row, colsCount)
		for j, elem := range row {
			parsed[header[j]] = elem
		}
		result = append(result, parsed)
	}
	return
}
