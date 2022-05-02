package log

import "encoding/json"

type SearchLog struct {
	StartTime int64  `json:"starttime" type:"int" need:"否" default:"0" information:"起始时间"`
	Page      int    `json:"page" type:"int" need:"是" default:"0" information:"第几页"`
	Limit     int    `json:"limit" type:"int" need:"是" default:"0" information:"每页显示的个数"`
	Count     int    `json:"count,empty" type:"int" need:"否" default:"0" information:"代码内部计算使用"`
	EndTime   int64  `json:"endtime" type:"int" need:"否" default:"0" information:"结束时间"`
	Classify  string `json:"classify" type:"string" need:"否" default:"" information:"根据类型过滤"`
	Ip        string `json:"ip" type:"string" need:"否" default:"" information:"根据ip过滤, 没启用"`
}

type Loglist struct {
	LogList []*LogRow `json:"loglist" type:"int" need:"否" default:"[]" information:"列表"`
	Count   int       `json:"count" type:"int" need:"是" default:"0" information:"总个数"`
	Page    int       `json:"page" type:"int" need:"是" default:"0" information:"页数"`
}

func (ll *Loglist) ErrorE(err error) []byte {
	send, _ := json.Marshal(ll)
	return send
}

func (ll *Loglist) NoRows() []byte {
	ll.LogList = make([]*LogRow, 0)
	send, _ := json.Marshal(ll)
	return send
}
func (ll *Loglist) Error(err string) []byte {
	send, _ := json.Marshal(ll)
	return send
}

type LogRow struct {
	Id       int    `json:"id"`
	Exectime int64  `json:"exectime"`
	Classify string `json:"classify"`
	Action   string `json:"action"`
	Ip       string `json:"ip"`
	UserName string `json:"username"`
}
