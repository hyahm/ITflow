package gapublic

import (
	"errors"
	"reflect"
)

//  拷贝值到另一个相同的结构中
func CopyStruct(dst interface{}, src interface{}) (err error) {
	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		errors.New("dst not a struct ")
	}
	if reflect.TypeOf(src).Kind() != reflect.Ptr {
		errors.New("src not a struct ")
	}
	l := reflect.TypeOf(src).Elem().NumField()

	for i := 0; i < l; i++ {
		// 获取结构体字段名
		name := reflect.TypeOf(src).Elem().Field(i).Name
		//  获取旧struct 的值
		v := reflect.ValueOf(src).Elem().Field(i)
		//  获取新struct 的字段
		x := reflect.ValueOf(dst).Elem().FieldByName(name)
		// 设置值
		x.Set(v)
	}
	return nil
}
