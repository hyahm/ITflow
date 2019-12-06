package buglog

import (
	"errors"
	"github.com/hyahm/golog"
	"itflow/bug/bugconfig"
	"time"
)

type Cs string

const (
	Login   = "login"
	User    = "user"
	Bug     = "bug"
	Version = "version"
	Project = "project"
	Env     = "env"
)

func (al *AddLog) insert(classify string, content string) error {
	//for

	//如果连接断开了，重新建立连接
	if al.Ip == "" {
		return errors.New("ip must be need")
	}
	// 如果ip在列表里面，直接跳过
	for _, v := range bugconfig.Exclude {
		if v == al.Ip {
			golog.Info("ip: %s is exclude", al.Ip)
			return nil
		}
	}
	_, err := bugconfig.Bug_Mysql.Insert("insert into log(exectime,classify,content,ip) values(?,?,?,?)",
		time.Now().Unix(), classify, content, al.Ip,
	)
	if err != nil {
		return err
	}
	return nil
	//Exclude
}
