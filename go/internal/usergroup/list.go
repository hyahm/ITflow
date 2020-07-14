package usergroup

type RespUserGroupList struct {
	UserGroupList []*RespUserGroup `json:"usergrouplist"`
	Code          int              `json:"code"`
	Msg           string           `json:"message,omitempty"`
}
