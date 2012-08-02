package matching

import(
	"appengine"
	"appengine/memcache"
	"strings"
)


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