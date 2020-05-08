package user

type ChangePasswod struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}
