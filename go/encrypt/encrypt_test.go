package encrypt

import "testing"

func TestSha1(t *testing.T) {
	p1 := PwdEncrypt("aaaa", "777777")
	t.Log(p1)
}
