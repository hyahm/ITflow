package midware

import (
	"encoding/json"
	"io/ioutil"
	"itflow/model/response"
	"net/http"

	"github.com/hyahm/xmux"
)

func JsonToStruct(w http.ResponseWriter, r *http.Request) bool {
	resp := &response.Response{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	err = json.Unmarshal(b, xmux.GetData(r).Data)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	return false
}
