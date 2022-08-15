package tree

import (
	"sort"
)

// ListToMap 将一个map类型的list转换成一个树状的map
//  - layers: 待转换的列表
func ListToMap(layers []map[string]interface{}) (result map[string]interface{}) {
	if len(layers) == 0 {
		return
	}
	// 先排个序
	sort.Slice(layers, func(i, j int) bool {
		return layers[i]["id"].(int) < layers[j]["id"].(int)
	})
	result = layers[0]
	for _, v := range layers[1:] {
		parentIdx := v["pid"].(int) - 1
		_, isExisted := layers[parentIdx]["children"]
		if !isExisted {
			layers[parentIdx]["children"] = []interface{}{v}
		} else {
			layers[parentIdx]["children"] = append(layers[parentIdx]["children"].([]interface{}), v)
		}
	}
	return
}
