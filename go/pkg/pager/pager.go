package pager

import (
	"github.com/hyahm/golog"
)

// 返回正确的起始值,和末尾数, 总数，页数， 索引为0 显示数
func GetPagingLimitAndPage(count int, page int, limit int) (int, int) {
	// 都小于1了
	if limit == 0 {
		return 0, 0
	}
	if page < 1 {
		return 1, count % limit
	}
	// 超出了，返回最大的页码
	if page*limit > count+limit {
		if count%limit == 0 {
			return ((count / limit) - 1) * limit, limit
		} else {
			return (count/limit + 1) * limit, count % limit
		}
	} else {
		// if count%limit == 0 {
		start := (page - 1) * limit
		golog.Info(start)

		if count-start < limit {
			return (page - 1) * limit, count - start
		} else {
			return (page - 1) * limit, limit
		}

	}
}
