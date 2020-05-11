package status

type Status struct {
	CheckStatus []string `json:"checkstatus"`
	Code        int      `json:"code"`
}
