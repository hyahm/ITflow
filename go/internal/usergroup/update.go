package usergroup

type RespUpdateUserGroup struct {
	Id    int64    `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
	Code  int      `json:"code"`
}
