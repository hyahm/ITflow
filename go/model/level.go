package model

import "itflow/cache"

type Table_level struct {
	Id   cache.LevelId `json:"id"`
	Name cache.Level   `json:"name"`
}

type Data_level struct {
	Id   cache.LevelId `json:"id"`
	Name cache.Level   `json:"name"`
	Code int           `json:"code"`
}

type List_levels struct {
	Levels []*Table_level `json:"levels"`
	Code   int            `json:"code"`
}

type Update_level struct {
	Id      cache.LevelId `json:"id"`
	Name    cache.Level   `json:"name"`
	OldName cache.Level   `json:"oldname"`
}
