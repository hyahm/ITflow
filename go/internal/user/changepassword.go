package user

type ResetPassword struct {
	Id       int    `json:"id"`
	Password string `json:"newpassword"`
}
