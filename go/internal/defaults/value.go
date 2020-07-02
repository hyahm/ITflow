package defaults

type DefaultValue struct {
	DefaultStatus string `json:"defaultstatus"`
	Important     string `json:"defaultimportant"`
	Level         string `json:"defaultlevel"`
	Code          int    `json:"code"`
}
