package card

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	. "library"
)


func Editcard(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	encodedKey := r.FormValue("key")
	decodedKey, err := datastore.DecodeKey(encodedKey)
	if err != nil {
		err.Error()
	}
	
	if r.FormValue("action") == "delete" {
		err = datastore.Delete(c, decodedKey)
		if err != nil {
			err.Error()
		}
		http.Redirect(w, r, "/cardlist", 303)
	}
	
	material := make(map[string]string)
	material["Key"] = encodedKey
	Output(w, "page/cardset/card/editcard.html", material)
}