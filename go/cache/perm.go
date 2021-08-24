package cache

// 增删改查顺序
var PermissionList = []string{"Read", "Create", "Update", "Delete"}

// 索引对应
var PermissionMap = map[string]int{
	"Read":   0,
	"Create": 1,
	"Update": 2,
	"Delete": 3,
}

// 获取所有用户对应权限

type PageInfo struct {
	Name string
	Info string
}
