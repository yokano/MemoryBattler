package ultrarich

import(
	"net/http"
	. "library"
)

func Ultrarich(w http.ResponseWriter, r *http.Request) {
	gamekey := r.FormValue("gamekey")
	Output(w, "game/ultrarich/ultrarich.html", make(map[string]string))
	if gamekey == "" {
		return
	}
}
