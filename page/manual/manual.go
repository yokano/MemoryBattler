package manual

import(
	"net/http"
	. "library"
)

func Manual(w http.ResponseWriter, r *http.Request) {
	Output(w, "page/manual/manual.html", make(map[string]string))
}