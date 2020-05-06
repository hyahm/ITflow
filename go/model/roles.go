package model

// 页面的权限表， 由开发者管理
type Roles struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`
}
