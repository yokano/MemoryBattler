package cardset

import(
	"net/http"
	. "library"
)

func Createcardset(w http.ResponseWriter, r *http.Request) {
	Output(w, "page/cardset/createcardset.html", make(map[string]string))
}