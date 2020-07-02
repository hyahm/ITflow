package handle

import (
	"fmt"
	"net/http"

	"github.com/hyahm/xmux"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(xmux.Var(r)["bb"])
}
