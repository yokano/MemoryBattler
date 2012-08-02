/*
	init.go
	yuta.okano@gmail.com
	
	クライアントからのリクエストURLに応じてハンドラを設定する
*/

package page

import(
	"appengine"
	"appengine/datastore"
	"net/http"
	
	. "page/cardset/card"
	. "page/cardset"
	. "page/game"
	. "page/manual"
	. "page/setting"
	. "game/ultrarich"
	. "library"
	. "library/matching"
)

func init() {
	// system
	http.HandleFunc("/", Top)
	http.HandleFunc("/top", Top)
	http.HandleFunc("/cardsetlist", Cardsetlist)
	http.HandleFunc("/cardlist", Cardlist)
	http.HandleFunc("/cardsetdetail", Cardsetdetail)
	http.HandleFunc("/createcardset", Createcardset)
	http.HandleFunc("/createcard", Createcard)
	http.HandleFunc("/creategame", Creategame)
	http.HandleFunc("/editcard", Editcard)
	http.HandleFunc("/play", Play)
	http.HandleFunc("/gamelist", Gamelist)
	http.HandleFunc("/manual", Manual)
	http.HandleFunc("/setting", Setting)
	
	// games
	http.HandleFunc("/ultrarich", Ultrarich)
	
	// debug
	http.HandleFunc("/setup", Setup)
	http.HandleFunc("/debug", Debug)
	http.HandleFunc("/matching", Debug)
}

func Debug(w http.ResponseWriter, r *http.Request) {
	Matching(w, r, "test", 2)
}

func Setup(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	rule := new(Rule)
	rule.Name = "ばばぬき"
	rule.Path = "/oldmaid"
	key := datastore.NewIncompleteKey(c, "Rule", nil)
	datastore.Put(c, key, rule)
	rule.Name = "七並べ"
	rule.Path = "/seven"
	key = datastore.NewIncompleteKey(c, "Rule", nil)
	datastore.Put(c, key, rule)
	rule.Name = "大富豪"
	rule.Path = "/ultrarich"
	key = datastore.NewIncompleteKey(c, "Rule", nil)
	datastore.Put(c, key, rule)
	c.Debugf("セットアップ完了")
	http.Redirect(w, r, "/top", 303)
}

