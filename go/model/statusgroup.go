package model

type StatusGroup struct {
	ID        int64   `json:"id" db:"id,default"`
	Name      string  `json:"name" db:"name"`
	StatusIDs []int64 `json:"sids" db:"sids"`
}
