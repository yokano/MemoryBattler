package cardset

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	"fmt"
	. "library"
)

func Cardsetlist(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	// add inputted cardset
	name := r.FormValue("name")
	if name != "" {
		cardset := new(Cardset)
		cardset.Name = name
		datastore.Put(c, datastore.NewIncompleteKey(c, "Cardset", nil), cardset)
	}
	
	// get all cardsets from datastore
	cardsets := new([]Cardset)
	q := datastore.NewQuery("Cardset")
	keys,err := q.GetAll(c, cardsets)
	if err != nil {
		err.Error()
	}

	// Create <li> of any games
	liTemplate := `<li><a href="/cardsetdetail?key=%s">%s</a></li>`
	cardset := new(Cardset)
	material := make(map[string]string)
	for i := 0; i < len(keys); i++ {
		datastore.Get(c, keys[i], cardset)
		material["Cardsets"] += fmt.Sprintf(liTemplate, keys[i].Encode(), cardset.Name)
	}
	Output(w, "page/cardset/cardsetlist.html", material)
}