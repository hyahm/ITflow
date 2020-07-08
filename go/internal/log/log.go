package log

type Search_log struct {
	StartTime int64  `json:"starttime" type:"int" need:"否" default:"0" information:"起始时间"`
	Page      int    `json:"page" type:"int" need:"是" default:"0" information:"第几页"`
	Limit     int    `json:"limit" type:"int" need:"是" default:"0" information:"每页显示的个数"`
	Count     int    `json:"count,empty" type:"int" need:"否" default:"0" information:"代码内部计算使用"`
	EndTime   int64  `json:"endtime" type:"int" need:"否" default:"0" information:"结束时间"`
	Classify  string `json:"classify" type:"string" need:"否" default:"" information:"根据类型过滤"`
	Ip        string `json:"ip" type:"string" need:"否" default:"" information:"根据ip过滤, 没启用"`
}

type Loglist struct {
	LogList []*LogRow `json:"loglist" type:"int" need:"否" default:"[]" information:"列表"`
	Code    int       `json:"code" type:"int" need:"是" default:"0" information:"状态码"`
	Count   int       `json:"count" type:"int" need:"是" default:"0" information:"总个数"`
	Page    int       `json:"page" type:"int" need:"是" default:"0" information:"页数"`
	Msg     string    `json:"msg, omitempty" type:"string" need:"否" default:"" information:"错误信息"`
}

type LogRow struct {
	Id       int    `json:"id"`
	Exectime int64  `json:"exectime"`
	Classify string `json:"classify"`
	Action   string `json:"action"`
	Ip       string `json:"ip"`
	UserName string `json:"username"`
}

func (sl *Search_log) GetPagingLimitAndPage() (int, int) {
	// 都小于1了
	if sl.Limit == 0 {
		return 0, 0
	}
	if sl.Page < 1 {
		sl.Page = 1
	}
	// 超出了，返回最大的页码
	if sl.Page*sl.Limit > sl.Count+sl.Limit {

		if sl.Count%sl.Limit == 0 {
			sl.Page = sl.Count / sl.Limit
			return ((sl.Count / sl.Limit) - 1) * sl.Limit, sl.Limit
		} else {
			sl.Page = sl.Count/sl.Limit + 1
			return (sl.Count/sl.Limit + 1) * sl.Limit, sl.Count % sl.Limit
		}
	} else {
		// if count%limit == 0 {

		start := (sl.Page - 1) * sl.Limit
		if sl.Count-start < sl.Limit {
			return start, sl.Count - start
		} else {
			return start, sl.Limit
		}

	}
}
