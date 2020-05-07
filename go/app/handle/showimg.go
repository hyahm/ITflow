package handle

import (
	"io/ioutil"
	"itflow/app/bugconfig"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ShowImg(w http.ResponseWriter, r *http.Request) {

	name := xmux.Var(r)["imgname"]
	golog.Info(name)
	golog.Infof("image name %s", filepath.Join(bugconfig.ImgDir, name))
	file, err := os.Open(filepath.Join(bugconfig.ImgDir, name))
	if err != nil {
		golog.Error(err)
		return
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)

	if err != nil {
		golog.Error(err)
		return
	}

	w.Write(buff)

}
