package handle

import (
	"encoding/json"
	"itflow/app/bugconfig"
	"net/http"
)

func GetRoleGroup(w http.ResponseWriter, r *http.Request) {

	rl := &getroles{}
	for _, v := range bugconfig.CacheRidGroup {
		rl.Rolelist = append(rl.Rolelist, v)
	}
	send, _ := json.Marshal(rl)
	w.Write(send)
	return

}
