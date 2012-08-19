package game

import(
	"net/http"
	"appengine"
	"appengine/datastore"
	"fmt"
	. "library"
	. "library/matching"
)

// Create and show gamelist page
func Gamelist(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	gamekey := r.FormValue("gamekey")
	
	if gamekey != "" {
		c.Debugf(gamekey)
		Matching(w, r, gamekey, 2)
	} else {
		// get all games from datastore
		q := datastore.NewQuery("Game")
		games := q.Run(c)
		gameNum,_ := q.Count(c)
		
		// Create <li> of any games
		liTemplate := 
			`
			<li>
				<a href="/gamelist?gamekey=%s" data-transition="slide">
					<div>%s</div>
					<div>ルール：%s</div>
					<div>カードセット：%s</div>
					<div>人数：2/4</div>
				</a>
			</li>
			`
		game := new(Game)
		material := make(map[string]string)
		cardset := new(Cardset)
		rule := new(Rule)
		var gamekey *datastore.Key
		for i := 0; i < gameNum; i++ {
			gamekey,_ = games.Next(game)
			datastore.Get(c, game.Cardset, cardset)
			datastore.Get(c, game.Rule, rule)
			material["GameList"] += fmt.Sprintf(liTemplate, gamekey.Encode(), game.Name, rule.Name, cardset.Name)
		}
		Output(w, "page/game/gamelist.html", material)
	}
	
}