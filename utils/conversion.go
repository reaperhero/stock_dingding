package utils

import (
	"fmt"
	"sort"
	"strings"
)

func SortMapWithKey(source map[interface{}]interface{}, asc bool) []interface{} {
	if len(source) <= 1 {
		return nil
	}
	var ks []interface{}
	for k, _ := range source {
		ks = append(ks, k)
	}
	switch ks[0].(type) {
	case string:
		sort.Slice(ks, func(i, j int) bool {
			return (strings.Compare(ks[i].(string), ks[j].(string)) > 0) == asc
		})
		return ks
	case int:
		sort.Slice(ks, func(i, j int) bool {
			return (ks[i].(int) > ks[j].(int)) == asc
		})
		return ks
	default:
		return nil
	}
}

func SortMapWithValue(source map[interface{}]interface{}, asc bool) []interface{} {
	if len(source) <= 1 {
		return nil
	}
	var (
		ks     []interface{}
		last   []int
		result []interface{}
	)
	for _, v := range source {
		ks = append(ks, v)
	}
	newKs := removeDuplicateElement(swapSliceToInt(ks))
	switch t := newKs.(type) {
	case []int:
		sort.Slice(t, func(i, j int) bool {
			return t[i]>t[j] == asc
		})
		last = t
	}
	for _, v := range last {
		for sK, sV := range source {
			if sV == v {
				result = append(result, sK)
			}
		}
	}
	return result
}

func SortMapWithValueSliceLen(source map[interface{}][]interface{}, asc bool) []interface{} {
	if len(source) <= 0 {
		return nil
	}
	var (
		ks     []int
		result []interface{}
	)
	for _, v := range source {
		ks = append(ks, len(v))
	}
	sort.Slice(ks, func(i, j int) bool {
		return ks[i] > ks[j] && asc
	})
	for _, k := range ks {
		for sK, sV := range source {
			if len(sV) == k {
				result = append(result, sK)
			}
		}
	}
	return result
}

func removeDuplicateElement(originals interface{}) interface{} {
	temp := map[string]struct{}{}
	switch slice := originals.(type) {
	case []string:
		var result []string
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result
	case []int64:
		var result []int64
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result
	case []int:
		var result []int
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result
	default:
		return nil
	}
}

func swapSliceToInt(source []interface{}) []int {
	var result []int
	for _, v := range source {
		result = append(result, v.(int))
	}
	return result
}
