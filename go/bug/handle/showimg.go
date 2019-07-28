package handle

import (
	"itflow/bug/bugconfig"
	"github.com/gorilla/mux"
	"github.com/hyahm/golog"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func ShowImg(w http.ResponseWriter, r *http.Request) {

	//headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {

		vars := mux.Vars(r)
		name := vars["imgname"]

		file, err := os.Open(path.Join(bugconfig.ImgDir, name))

		if err != nil {
			golog.Error(err.Error())
			return
		}
		defer file.Close()
		buff, err := ioutil.ReadAll(file)

		if err != nil {
			golog.Error(err.Error())
			return
		}

		w.Write(buff)

	}

}
