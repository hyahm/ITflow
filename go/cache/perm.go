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
	Page string
	Info string
}

var PageAndRole = []PageInfo{
	{
		Page: "env",
		Info: "环境页面",
	},
	{
		Page: "important",
		Info: "重要性页面",
	},
	{
		Page: "level",
		Info: "优先级别页面",
	},
	{
		Page: "position",
		Info: "职位页面",
	},
	{
		Page: "project",
		Info: "项目页面",
	},
	{
		Page: "status",
		Info: "bug状态流程页面",
	},
	{
		Page: "statusgroup",
		Info: "状态组页面",
	},
	{
		Page: "version",
		Info: "版本页面",
	},
}
