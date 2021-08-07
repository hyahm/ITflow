package status

type ChangeStatus struct {
	Status []string `json:"status"`
	Code   int      `json:"code"`
}
