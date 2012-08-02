package cardset

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	. "library"
)


func Cardsetdetail(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	// get key
	encodedKey := r.FormValue("key")
	decodedKey, err := datastore.DecodeKey(encodedKey)
	if err != nil {
		err.Error()
	}
	
	// delete mode
	if r.FormValue("action") == "delete" {
		err = datastore.Delete(c, decodedKey)
		if err != nil {
			err.Error()
		}
		http.Redirect(w, r, "/cardsetlist", 303)
	}
		
	// get cardset
	cardset := new(Cardset)
	err = datastore.Get(c, decodedKey, cardset)
	if err != nil {
		err.Error()
	}
	
	// output html
	material := make(map[string]string)
	material["Name"] = cardset.Name
	material["Key"] = encodedKey
	Output(w, "page/cardset/cardsetdetail.html", material)
}