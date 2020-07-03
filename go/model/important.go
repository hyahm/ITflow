package model

import (
	"itflow/cache"
)

type Importants struct {
	Id   cache.ImportantId `json:"id"`
	Name cache.Important   `json:"name"`
}

type Data_importants struct {
	Id   cache.ImportantId `json:"id"`
	Name cache.Important   `json:"name"`
	Code int               `json:"code"`
}

type List_importants struct {
	ImportantList []*Importants `json:"importantlist"`
	Code          int           `json:"code"`
}
