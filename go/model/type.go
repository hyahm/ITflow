package model

type Data_types struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name"`
	Types    int64      `json:"checktype"`
	Opts     []*Options `json:"opts"`
	Listtype string     `json:"listtype"`
	Default  string     `json:"default"`
	Code     int        `json:"code"`
}

type List_types struct {
	List []*Data_types `json:"list"`
	Code int           `json:"code"`
}

type Send_types struct {
	Id   int64      `json:"id"`
	Code int        `json:"code"`
	Opts []*Options `json:"opts"`
}

type Options struct {
	Id      int64  `json:"id"`
	Info    string `json:"info"`
	Name    string `json:"name"`
	Need    string `json:"need"`
	Type    string `json:"type"`
	Default string `json:"default"`
}
