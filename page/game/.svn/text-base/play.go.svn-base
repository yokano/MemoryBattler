package game

import(
	"appengine"
	"appengine/datastore"
	"net/http"
	. "library"
)

func Play(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	name := r.FormValue("name")
	cardsetKey := r.FormValue("cardset")
	ruleKey := r.FormValue("rule")
	if name != "" && cardsetKey != "" && ruleKey != "" {
		
		// datastoreへゲームを登録
		game := new(Game)
		game.Name = name
		game.Cardset,_ = datastore.DecodeKey(cardsetKey)
		game.Rule,_ = datastore.DecodeKey(ruleKey)
		gamekey,_ := datastore.Put(c, datastore.NewIncompleteKey(c, "Game", nil), game)
		
		// ゲーム画面へ移動
		rule := new(Rule)
		datastore.Get(c, game.Rule, rule)
		path := rule.Path + "?gamekey=" + gamekey.Encode()
		http.Redirect(w, r, path, 303)
	}
}