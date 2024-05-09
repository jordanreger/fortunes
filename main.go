package fortunes

import (
	_ "embed"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"jordanreger.com/web/util"
)

//go:embed fortunes
var fortunes string

func Fortune(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	var path string
	path = util.FormatPath(url.Path)

	last := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]

	f := strings.Split(string(fortunes), "\n")

	id, _ := strconv.ParseInt(last, 10, 64)
	if id >= 1 && id <= 4288 {
		fmt.Fprint(w, f[id-1])
		return
	} else {
		fmt.Fprint(w, f[rand.Intn(len(f))])
		return
	}
}
