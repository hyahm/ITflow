package search

import "testing"

func TestMybugs(t *testing.T) {
	mybugs := &ReqMyBugFilter{
		page:    1,
		Limit:   10,
		Level:   "",
		Project: "",
		Title:   "",
		// ShowsStatus cache.StatusList // 这个应该从数据库获取
	}
	var uid int64 = 1
	err := mybugs.GetUsefulCondition(uid)
	if err != nil {
		t.Log(err)
	}

}
