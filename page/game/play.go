package game

import(
	"appengine"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Debugf("play")
	
	// ゲーム画面へ移動
//	rule := new(Rule)
//	datastore.Get(c, game.Rule, rule)
//	path := rule.Path + "?gamekey=" + gamekey.Encode()
//	http.Redirect(w, r, path, 303)
}