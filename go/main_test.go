package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type Jsonstr string

func (js Jsonstr) Json() (string, error) {

	b := make([]byte, 2048)
	buf := bytes.NewBuffer(b)
	err := json.Indent(buf, []byte(js), "", "    ")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func TestJson(t *testing.T) {
	fmt.Println(8&7 == 0)
	// var a Jsonstr = `{"loglist":[{"id":26,"exectime":1588840365,"classify":"login","content":"","ip":"127.0.0.1"},{"id":25,"exectime":1588840233,"classify":"login","content":"","ip":"127.0.0.1"},{"id":24,"exectime":1588837232,"classify":"login","content":"","ip":"127.0.0.1"},{"id":23,"exectime":1588837002,"classify":"login","content":"","ip":"127.0.0.1"},{"id":22,"exectime":1588833133,"classify":"login","content":"","ip":"127.0.0.1"},{"id":21,"exectime":1588833047,"classify":"login","content":"","ip":"127.0.0.1"}],"code":0, "count":100}`
	// j, err := a.Json()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf(j)
}
