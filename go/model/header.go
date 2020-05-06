package model

type Table_header struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Hhids  string `json:"hhids"`
	Remark string `json:"remark"`
}

type Data_header struct {
	Id     int64               `json:"id"`
	Name   string              `json:"name"`
	Hhids  []*Table_headerlist `json:"hhids"`
	Remark string              `json:"remark"`
	Code   int                 `json:"code"`
}

type Table_headerlist struct {
	Id    int64  `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type List_headers struct {
	Headers []*Data_header `json:"headers"`
	Code    int            `json:"code"`
}

type Update_header struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	OldName string `json:"oldname"`
}
