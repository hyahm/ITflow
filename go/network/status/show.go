package status

type Status struct {
	CheckStatus []string `json:"checkstatus"`
	Code        int      `json:"code"`
}

type ChangeStatus struct {
	Status []string `json:"status"`
	Code   int      `json:"code"`
}
