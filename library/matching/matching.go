package matching

import (
	. "library"
	"net/http"
	"encoding/json"
	"appengine"
	"appengine/channel"
)

// Config
type Config struct {
	maxplayer uint8
	gamekey string
}


// Client
type Client struct {
	name string
	id string
	token string
	action string
}
func newClient(context appengine.Context, r *http.Request) *Client {
	c := new(Client)
	c.name = r.FormValue("name")
	c.action = r.FormValue("action")
	
	// ユーザIDが設定されていなければ新しく設定する
	c.id = r.FormValue("id")
	if c.id == "" {
		c.id = CreateId()
	}
	
	var err error
	c.token,err = channel.Create(context, c.id)
	Check(context, err)
	return c
}

func Matching(w http.ResponseWriter, r *http.Request, gamekey string, maxplayer uint8) {
	c := appengine.NewContext(r)
	client := newClient(c, r)
	config := Config{maxplayer, gamekey}
	memcache := newMemcache(c, &config)
	
	actions := make(map[string]func())
	
	// メモリにプレイヤーを追加してチャネルトークンを返す
	actions["join"] = func() {
		memcache.AddPlayer(c, client.id, client.name)
		message := map[string]string{"token":client.token, "id":client.id}
		result,err := json.Marshal(message)
		Check(c, err)
		w.Write(result)
	}
	
	// メモリからプレイヤーを削除
	actions["leave"] = func() {
		memcache.DeletePlayer(c, client.id)
	}
	
	// ユーザ一覧を返す
	actions["get"] = func() {
		w.Write(memcache.GetPlayersJson(c))
	}
	
	if client.action == "" {
		Output(w, "library/matching/matching.html", make(map[string]string))
	} else {
		actions[client.action]()
	}
}
