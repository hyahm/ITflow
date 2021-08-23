package rolegroup

import (
	"itflow/model"
)

type PermRole struct {
	Add    bool   `json:"add"`
	Remove bool   `json:"remove"`
	Select bool   `json:"select"`
	Update bool   `json:"update"`
	Info   string `json:"info"`
	Name   string `json:"name"`
}

type ReqRoleGroup struct {
	Id    int64        `json:"id" type:"int" need:"是" default:"" information:"id"`
	Name  string       `json:"name" type:"string" need:"是" default:"" information:"角色组名"`
	Perms []model.Perm `json:"rolelist" type:"array" need:"是" default:"" information:"角色组成员"`
}

type RoleGroupList struct {
	DataList []*ReqRoleGroup `json:"datalist" type:"array" need:"否" default:"" information:"角色组列表"`
	Code     int             `json:"code" type:"int" need:"是" default:"" information:"状态码"`
}

type Updata_role struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Oldname string `json:"oldname"`
}

func (reqrg *ReqRoleGroup) Add(uid int64) []byte {
	return nil
	// 添加角色组权限
	// errrocode := &response.Response{}
	// // 必须是管理员, 权限问题后面全部移动到中间件
	// if uid != cache.SUPERID {
	// 	golog.Error("no perm")
	// 	return errrocode.Error("no perm")
	// }
	// if reqrg.Name == "" {
	// 	return errrocode.Error("名称不能为空")
	// }
	// // 过滤
	// pernids := make([]string, 0)
	// golog.Info("info: ", cache.CacheRoleRid)
	// // 确保每个info 没漏掉， 应该要加个判断的
	// for _, v := range reqrg.RoleList {
	// 	// 不管有没有效， 全部插入
	// 	if !v.Select {
	// 		v.Remove = false
	// 		v.Update = false
	// 		v.Add = false
	// 	}
	// 	golog.Info("info: ", v.Info)

	// 	if rid, ok := cache.CacheRoleRid[v.Name]; ok {
	// 		perm := &model.Perm{
	// 			Find:     v.Select,
	// 			Remove:   v.Remove,
	// 			Revise:   v.Update,
	// 			Increase: v.Add,
	// 			Rid:      rid,
	// 		}
	// 		golog.Info(*perm)
	// 		// 数据插入到perm表
	// 		err := perm.Insert()
	// 		if err != nil {
	// 			golog.Info(err)
	// 			return errrocode.ErrorE(err)
	// 		}
	// 		pernids = append(pernids, strconv.FormatInt(perm.Id, 10))
	// 	} else {
	// 		golog.Info(v.Name)
	// 	}

	// }
	// // 数据插入到rolegroup
	// rg := &model.RoleGroup{
	// 	Name:    reqrg.Name,
	// 	Permids: strings.Join(pernids, ","),
	// }
	// err := rg.Insert()
	// if err != nil {
	// 	golog.Error(err)
	// 	return errrocode.ErrorE(err)
	// }
	// errrocode.ID = rg.ID
	// return errrocode.Success()
}

func (reqrg *ReqRoleGroup) Update(uid int64) []byte {
	// errrocode := &response.Response{}
	// rg := &model.RoleGroup{
	// 	ID:   reqrg.Id,
	// 	Name: reqrg.Name,
	// }
	// if uid != cache.SUPERID {
	// 	golog.Error("no perm")
	// 	return errrocode.Error("no perm")
	// }
	// // 修改rolegroup表的值， 主要是name， 其他的都是固定的
	// err := rg.Update()
	// if err != nil {
	// 	golog.Info(err)
	// 	return errrocode.ErrorE(err)
	// }

	// // 更新 perm 表的值
	// for _, v := range reqrg.RoleList {
	// 	golog.Infof("%+v", v)
	// 	// 不管有没有效， 全部插入
	// 	if !v.Select {
	// 		v.Remove = false
	// 		v.Update = false
	// 		v.Add = false
	// 	}
	// 	if rid, ok := cache.CacheRoleRid[v.Name]; ok {
	// 		perm := &model.Perm{
	// 			Find:     v.Select,
	// 			Rid:      rid,
	// 			Remove:   v.Remove,
	// 			Revise:   v.Update,
	// 			Increase: v.Add,
	// 		}
	// 		if rid == 2 {
	// 			golog.Info(perm.Increase)
	// 		}
	// 		// 根据范围内的唯一性 rid 确地唯一性
	// 		err := perm.Update(rg.Permids)
	// 		if err != nil {
	// 			golog.Info(err)
	// 			return errrocode.ErrorE(err)
	// 		}

	// 	} else {
	// 		golog.Info(v.Name)
	// 	}

	// }
	return nil
	// return errrocode.Success()
}
