package matching

import (
	. "library"
	"net/http"
	"encoding/json"
	"appengine"
	"appengine/memcache"
	"appengine/channel"
	"strings"
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
	if err != nil {
		context.Debugf(err.Error())
	}
	return c
}


// Memacache
type Memcache struct {
	config *Config
}
func newMemcache(c appengine.Context, config *Config) *Memcache {
	m := new(Memcache)
	m.config = config
	
	// memcache がなかったら作る
	_,err := memcache.Get(c, m.config.gamekey)
	if err != nil {
		item := memcache.Item{
			Key: m.config.gamekey,
			Value: []byte(""),
		}
		memcache.Set(c, &item)
	}
	
	return m
}
func (m *Memcache) AddPlayer(c appengine.Context, id string) {
	oldItem,err := memcache.Get(c, m.config.gamekey)
	if err != nil {
		c.Debugf(err.Error())
	}
	
	oldValue := (string)(oldItem.Value)
	if oldValue != "" {
		id = oldValue + "," + id
	}
	
	newItem := memcache.Item {
		Key: m.config.gamekey,
		Value: []byte(id),
	}
	memcache.Set(c, &newItem)
}
func (m *Memcache) DeletePlayer(c appengine.Context, id string) {
	oldPlayers := m.GetPlayers(c)
	newPlayers := []string{}
	
	for i := 0; i < len(oldPlayers); i++ {
		if oldPlayers[i] != id {
			newPlayers = append(newPlayers, oldPlayers[i])
		}
	}
	
	deleted := memcache.Item {
		Key: m.config.gamekey,
		Value: []byte(strings.Join(newPlayers, ",")),
	}
	memcache.Set(c, &deleted)
}
func (m *Memcache) GetPlayers(c appengine.Context) []string {
	item,err := memcache.Get(c, m.config.gamekey)
	if err != nil {
		c.Debugf(err.Error())
	}
	
	value := (string)(item.Value)
	result := []string{}
	
	if value != "" {
		result = strings.Split(value, ",")
	}
	
	return result
}

func Matching(w http.ResponseWriter, r *http.Request, gamekey string, maxplayer uint8) {
	c := appengine.NewContext(r)
	client := newClient(c, r)
	config := Config{maxplayer, gamekey}
	memcache := newMemcache(c, &config)
	
	actions := make(map[string]func())
	
	// メモリにプレイヤーを追加してチャネルトークンを返す
	actions["join"] = func() {
		memcache.AddPlayer(c, client.id)
		message := map[string]string{"token":client.token, "id":client.id}
		result,err := json.Marshal(message)
		if err != nil {
			c.Debugf(err.Error())
		}
		w.Write(result)
	}
	
	// メモリからプレイヤーを削除
	actions["leave"] = func() {
		memcache.DeletePlayer(c, client.id)
	}
	
	if client.action == "" {
		Output(w, "library/matching/matching.html", make(map[string]string))
	} else {
		actions[client.action]()
	}
}
