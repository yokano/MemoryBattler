package matching

import(
	"appengine"
	"appengine/memcache"
	"encoding/json"
	. "library"
)

type Memcache struct {
	config *Config
}

// コンストラクタ
func newMemcache(c appengine.Context, config *Config) *Memcache {
	m := new(Memcache)
	m.config = config
	_,err := memcache.Get(c, m.config.Gamekey)
	if err != nil {
		memcache.Set(c, &memcache.Item{ Key: m.config.Gamekey, Value: []byte("") })
	}
	return m
}

/*
	AddPlayer(c, id, name)
	メモリにプレイヤーを追加する
*/
func (m *Memcache) AddPlayer(c appengine.Context, id string, name string) {
	prevPlayers := m.GetPlayers(c)
	
	nextPlayers := []map[string]string{}
	for i := range prevPlayers {
		nextPlayers = append(nextPlayers, prevPlayers[i])
	}
	nextPlayers = append(nextPlayers, map[string]string{"id":id, "name":name})
	
	m.SetPlayers(c, nextPlayers)
}

/*
	DeletePlayer(c, id)
	メモリからプレイヤーを削除する
*/
func (m *Memcache) DeletePlayer(c appengine.Context, id string) {
	prevPlayers := m.GetPlayers(c)
	nextPlayers := []map[string]string{}
	for i := range prevPlayers {
		if(prevPlayers[i]["id"] != id) {
			nextPlayers = append(nextPlayers, prevPlayers[i])
		}
	}
	m.SetPlayers(c, nextPlayers)
	c.Debugf("DeletePlayer")
}

/*
	GetPlayers(c)
	メモリ上のプレイヤー一覧を取得する
*/
func (m *Memcache) GetPlayers(c appengine.Context) []map[string]string {
	memory,err := memcache.Get(c, m.config.Gamekey)
	Check(c, err)
	
	players := []map[string]string{}
	if( (string)(memory.Value) != "") {
		err = json.Unmarshal(memory.Value, &players)
		Check(c, err)
	}
	
	return players
}

/*
	GetPlayersJson(c)
	メモリ上のプレイヤー一覧をJSON形式で取得する
*/
func (m *Memcache) GetPlayersJson(c appengine.Context) []byte {
	memory,err := memcache.Get(c, m.config.Gamekey)
	Check(c, err)
	return memory.Value
}

/*
	SetPlayers(c)
	メモリ上にプレイヤーをセットする
*/
func (m *Memcache) SetPlayers(c appengine.Context, players []map[string]string) {
	value,err := json.Marshal(players)
	Check(c, err)
	memcache.Set(c, &memcache.Item{Key: m.config.Gamekey, Value: value})
}