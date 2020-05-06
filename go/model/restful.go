package model

type Table_restful struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	OwnerId  int    `json:"ownerid"`
	Perm     bool   `json:"perm"`
	Readuser bool   `json:"readuser"`
	Edituser bool   `json:"edituser"`
	Rid      int    `json:"rid"`
	Eid      int    `json:"eid"`
}

type Data_restful struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Auth     bool   `json:"auth"`
	Readuser bool   `json:"readuser"`
	Edituser bool   `json:"edituser"`
	ReadName string `json:"readname"`
	EditName string `json:"editname"`
	Resp     string `json:"resp"`
}

type List_restful struct {
	Code int             `json:"code"`
	List []*Data_restful `json:"list"`
}
type Table_apilist struct {
	Id          int    `json:"id"`
	Pid         int    `json:"pid"`
	Url         string `json:"url"`
	Information string `json:"information"`
	Opts        string `json:"opts"`
	Methods     string `json:"methods"`
	Result      string `json:"result"`
	Name        string `json:"name"`
}

type Get_apilist struct {
	Id          int           `json:"id"`
	Pid         int           `json:"pid"`
	Url         string        `json:"url"`
	Information string        `json:"information"`
	Opts        []*Table_opts `json:"opts"`
	Header      string        `json:"header"`
	Methods     []string      `json:"methods"`
	Result      string        `json:"result"`
	Name        string        `json:"name"`
	CallType    string        `json:"calltype"`
	Resp        string        `json:"resp"`
	Code        int           `json:"code"`
}

type Show_apilist struct {
	Id          int                 `json:"id"`
	Pid         int                 `json:"pid"`
	Url         string              `json:"url"`
	Information string              `json:"information"`
	Opts        []*Table_opts       `json:"opts"`
	Header      []*Table_headerlist `json:"header"`
	Methods     []string            `json:"methods"`
	Result      string              `json:"result"`
	Name        string              `json:"name"`
	CallType    string              `json:"calltype"`
	Code        int                 `json:"code"`
	Remark      string              `json:"remark"`
	Resp        string              `json:"resp"`
}

type One_apilist struct {
	Id          int           `json:"id"`
	Pid         int           `json:"pid"`
	Url         string        `json:"url"`
	Information string        `json:"information"`
	Opts        []*Table_opts `json:"opts"`
	Header      string        `json:"header"`
	Methods     []string      `json:"methods"`
	Result      string        `json:"result"`
	Name        string        `json:"name"`
	CallType    string        `json:"calltype"`
	Code        int           `json:"code"`
	Remark      string        `json:"remark"`
	Resp        string        `json:"resp"`
}

type Table_opts struct {
	Name    string `json:"name"`
	Info    string `json:"info"`
	Id      int    `json:"id"`
	Need    string `json:"need"`
	Type    string `json:"type"`
	Default string `json:"default"`
}

type List_api struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Data_api struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pid  int    `json:"pid"`
}

type Send_Types struct {
	Id    int64    `json:"id"`
	Code  int      `json:"code"`
	Types []string `json:"types"`
}
