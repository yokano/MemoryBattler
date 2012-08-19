package game

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	. "library"
	. "library/matching"
)

func Creategame(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	mode := r.FormValue("mode")
	
	if mode == "create" {
		create(c, w, r)
	} else if mode == "input" {
		input(c, w, r)
	}
}

// 新規ゲームを datastore へ追加
func create(c appengine.Context, w http.ResponseWriter, r *http.Request) {
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
		
		// マッチング画面を表示　マッチング完了後は play.go へ
		Matching(w, r, gamekey.Encode(), 2)
	}
}

// 新規ゲームの情報入力
func input(c appengine.Context, w http.ResponseWriter, r *http.Request) {
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