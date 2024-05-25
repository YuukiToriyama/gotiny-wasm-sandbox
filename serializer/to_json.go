package serializer

import (
	"github.com/YuukiToriyama/gotiny-wasm-sandbox/types"
	"sort"
)

func ToJson(rows types.Rows) (result string) {
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
