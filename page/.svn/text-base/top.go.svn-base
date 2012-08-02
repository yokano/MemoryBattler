/*
    top.go
    yuta.okano@gmail.com

    トップページの表示
*/

package page

import(
	"net/http"
	. "library"
)

func Top(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["Theme"] = "d"
	Output(w, "page/top.html", m)
}