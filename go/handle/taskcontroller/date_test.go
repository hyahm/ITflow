package taskcontroller

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	t.Log(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
}
