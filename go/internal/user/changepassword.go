package user

type ChangePasswod struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

type ResetPassword struct {
	Id       int    `json:"id"`
	Password string `json:"newpassword"`
}
