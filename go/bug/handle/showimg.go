package handle

import (
	"io/ioutil"
	"itflow/bug/bugconfig"
	"net/http"
	"os"
	"path"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ShowImg(w http.ResponseWriter, r *http.Request) {

	name := xmux.Var[r.URL.Path]["imgname"]
	golog.Info("image name %s", name)
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
