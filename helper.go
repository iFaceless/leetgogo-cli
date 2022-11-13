package leetcode

import "encoding/json"

func JSONMarshalMust(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
