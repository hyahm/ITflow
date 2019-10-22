package buglog

import (
	"fmt"
)

type AddLog struct {
	Ip       string
	Classify string
}

var Flog *AddLog

func init() {
	Flog = &AddLog{}
}

func (al *AddLog) Login(format string, args ...interface{}) error {
	return al.insert(al.Classify, fmt.Sprintf(format, args...))
}

func (al *AddLog) Add(args ...interface{}) error {
	return al.insert(al.Classify, fmt.Sprintf("add"+al.Classify+":operator:%v, id: %v, name: %v", args...))
}

func (al *AddLog) Del(args ...interface{}) error {
	return al.insert(al.Classify, fmt.Sprintf("delete"+al.Classify+": operator:%v, id: %v", args...))
}

func (al *AddLog) Update(args ...interface{}) error {
	return al.insert(al.Classify, fmt.Sprintf("update"+al.Classify+": operator:%v, id: %v, change to %v", args...))
}

type Log interface {
	Add() error
	Del() error
	Update() error
}
