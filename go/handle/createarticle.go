package handle

import (
	"net/http"
	//"strings"
)

type projectList struct {
	ProjectList []string `json:"projectlist"`
	Code        int      `json:"code"`
}

func GetProject(w http.ResponseWriter, r *http.Request) {

	// pl := &projectList{}

	// // for _, v := range cache.CachePidProject {
	// // 	pl.ProjectList = append(pl.ProjectList, v.ToString())
	// // }
	// send, _ := json.Marshal(pl)
	// w.Write(send)
	// return

}

// func GetMyProject(w http.ResponseWriter, r *http.Request) {
// 	myproject := &project.MyProject{
// 		Name: make([]string, 0),
// 	}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	w.Write(myproject.Get(uid))
// 	return

// }
