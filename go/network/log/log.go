package log

type Search_log struct {
	Id        int64  `json:"id"`
	StartTime int64  `json:"starttime"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	EndTime   int64  `json:"endtime"`
	Classify  string `json:"classify"`
	Content   string `json:"content"`
	Ip        string `json:"ip"`
}
