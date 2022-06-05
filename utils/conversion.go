package utils

import (
	"fmt"
	"reflect"
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


func StructAssignAonsistent(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)
	// 简单判断下
	if at.Kind() != reflect.Ptr {
		return fmt.Errorf("a must be a struct pointer")
	}
	av = reflect.ValueOf(av.Interface())
	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}
	if len(_fields) == 0 {
		return nil
	}
	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)
		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			continue
		}
	}
	return
}

func StructAssignSimilar(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}