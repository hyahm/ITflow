package user

type GetAddUser struct {
	Nickname    string   `json:"nickname"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Role        []string `json:"role"`
	RealName    string   `json:"realname"`
	RoleGroup   string   `json:"rolegroup"`
	StatusGroup string   `json:"statusgroup"`
	Position    string   `json:"position"` // 普通用户就是真，管理员就假
}
