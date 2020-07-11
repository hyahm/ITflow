package dashboard

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/model"
	"log"
	"time"
)

type BugCount struct {
	Created   []int  `json:"created"`
	Completed []int  `json:"completed"`
	Code      int    `json:"code"`
	Msg       string `json:"message"`
}

const SHOWDAY = 7

func GetCount() []byte {
	start, end := getTime()
	bc := &BugCount{
		Created:   make([]int, SHOWDAY),
		Completed: make([]int, SHOWDAY),
	}
	count := 0
	for i := SHOWDAY - 1; i >= 0; i-- {
		if cache.DefaultCreateSid > 0 {
			bc.Created[count], _ = model.GetCreatedCountByTime(start[i], end[i], cache.DefaultCreateSid.ToInt64())
		}
		if cache.DefaultCompleteSid > 0 {
			bc.Completed[count], _ = model.GetCompletedCountByTime(start[i], end[i], cache.DefaultCompleteSid.ToInt64())
		}
		count++
	}
	send, _ := json.Marshal(bc)
	return send
}

func getTime() ([]int64, []int64) {
	start := make([]int64, SHOWDAY)
	end := make([]int64, SHOWDAY)
	now := time.Now()
	start[0] = zeroTimeUnix(now)
	end[0] = now.Unix()
	for i := 1; i < SHOWDAY; i++ {
		start[i] = zeroTimeUnix(now.AddDate(0, 0, -1))
		end[i] = start[i-1] - 1
	}

	return start, end
}

func zeroTimeUnix(now time.Time) int64 {
	month := int(now.Month())
	m := fmt.Sprintf("%d", month)
	if month/10 < 1 {
		m = fmt.Sprintf("0%d", month)
	}
	day := now.Day()
	d := fmt.Sprintf("%d", day)
	if day/10 < 1 {
		d = fmt.Sprintf("0%d", day)
	}

	zeroTime := fmt.Sprintf("%d-%s-%s 00:00:00", now.Year(), m, d)
	t, err := time.Parse("2006-01-02 15:04:05", zeroTime)
	if err != nil {
		log.Fatal(err)
	}

	return t.Unix() - 8*60*60
}
