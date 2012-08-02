package card

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	"fmt"
	. "library"
)

func Cardlist(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	// add inputted game
	name := r.FormValue("name")
	if name != "" {
		card := new(Card)
		card.Name = name
		datastore.Put(c, datastore.NewIncompleteKey(c, "Card", nil), card)
		http.Redirect(w, r, "/cardlist", 303)
	}
	
	// get all cards from datastore
	cards := new([]Card)
	q := datastore.NewQuery("Card")
	keys,err := q.GetAll(c, cards)
	if err != nil {
		err.Error()
	}
	
	// Create <li> of any games
	material := make(map[string]string)
	liTemplate := `<li><a href="/editcard?key=%s" data-transition="slide">%s</a></li>`
	card := new(Card)
	for i := 0; i < len(keys); i++ {
		datastore.Get(c, keys[i], card)
		material["Cards"] += fmt.Sprintf(liTemplate, keys[i].Encode(), card.Name)
	}
	Output(w, "page/cardset/card/cardlist.html", material)
}