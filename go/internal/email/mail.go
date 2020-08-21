package email

type Email struct {
	Host      string `json:"host"`
	Enable    bool   `json:"enable"`
	Id        int64  `json:"id"`
	Port      int    `json:"port"`
	EmailAddr string `json:"email"`
	Password  string `json:"password"`
	To        string `json:"to"`
	Code      int    `json:"code"`
	Msg       int    `json:"msg,omitempty"`
}
