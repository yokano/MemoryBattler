package game

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	. "library"
)

func Creategame(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
		
	// create html template
	material := make(map[string]string)

	// create cardset select
	q := datastore.NewQuery("Cardset")
	cardsets := new([]Cardset)
	cardsetKeys,err := q.GetAll(c, cardsets)
	cardset := new(Cardset)
	if err != nil {
		err.Error()
	}
	for i := 0; i < len(cardsetKeys); i++ {
		err = datastore.Get(c, cardsetKeys[i], cardset)
		if err != nil {
			err.Error()
		}
		material["Cardsets"] += "<option value=\"" + cardsetKeys[i].Encode() + "\">" + cardset.Name + "</option>"
	}
	
	// create html of rule select
	q = datastore.NewQuery("Rule")
	rules := new([]Rule)
	rule := new(Rule)
	ruleKeys,err := q.GetAll(c, rules)
	if err != nil {
		err.Error()
	}
	for i := 0; i < len(ruleKeys); i++ {
		err = datastore.Get(c, ruleKeys[i], rule)
		if err != nil {
			err.Error()
		}
		material["Rules"] += "<option value=\"" + ruleKeys[i].Encode() + "\">" + rule.Name + "</option>"
	}
	
	Output(w, "page/game/creategame.html", material)
}