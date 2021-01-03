package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"itflow/cache"
	"strings"
	"time"

	"github.com/hyahm/golog"
)

const header = `{'typ': 'JWT', 
'alg': 'HS256'
}`

type Token struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Exp      int64  `json:"exp"`
}

func MakeJwt(id int64, nickname string) string {

	payload := fmt.Sprintf(`
	{
		"id": %d,
		"nickname": "%s",
		"exp": %d
	}
	`, id, nickname, time.Now().Add(time.Duration(cache.Expirontion)).Unix())

	s := base64.StdEncoding.EncodeToString([]byte(header))
	p := base64.StdEncoding.EncodeToString([]byte(payload))
	pre := s + "." + p
	token := pre + "." + getHc(pre)
	return token
}

func getHc(b string) string {
	h := hmac.New(sha256.New, []byte(cache.Salt))
	io.WriteString(h, b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (t *Token) CheckJwt(jwt string) bool {
	js := strings.Split(jwt, ".")

	b, err := base64.StdEncoding.DecodeString(js[1])
	if err != nil {
		golog.Error(err)
		return false
	}

	err = json.NewDecoder(bytes.NewReader(b)).Decode(t)
	if err != nil {
		golog.Error(err)
		return false
	}

	if t.Exp < time.Now().Unix() {
		golog.Error("token 过期")
		return false
	}
	// 检查过期时间
	pre := js[0] + "." + js[1]
	return getHc(pre) == js[2]
}
