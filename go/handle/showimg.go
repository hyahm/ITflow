package handle

import (
	"io/ioutil"
	"itflow/cache"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ShowImg(w http.ResponseWriter, r *http.Request) {

	name := xmux.Var(r)["imgname"]
	file, err := os.Open(filepath.Join(cache.ImgDir, name))
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
