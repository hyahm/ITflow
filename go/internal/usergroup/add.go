package usergroup

type RespUserGroup struct {
	Id    int64    `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
