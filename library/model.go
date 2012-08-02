/*
	model.go
	データモデルの定義
*/
package library
import "appengine/datastore"

const(
	SETUP = 1
	WAIT = 2
	START = 3
)

type Game struct {
	Name string
	Cardset *datastore.Key
	Rule *datastore.Key
	MaxPlayer int
	State int
}

type Cardset struct {
	Name string
}

type Card struct {
	Name string
}

type Rule struct {
	Name string
	Path string
}
