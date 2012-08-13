/*
	tools.go 2012/6/29
	yuta.okano@gmail.com
	
	汎用的に使用する関数群
*/

package library

import (
	"appengine"
	"text/template"
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"appengine/memcache"
	"appengine/channel"
)

/*
	関数 Output(path, material)
	- HTMLテンプレートにデータを置換してクライアントへ出力する

	引数
	- path : HTMLテンプレートのファイル名(page/*** の *** 部分)
	- material : 置換するデータを含むマップ
	
	戻り値
	- なし : クライアントの画面にHTMLを出力
*/
func Output(w http.ResponseWriter, path string, material interface{}) {
	tmpl,_ := template.ParseFiles(path)
	tmpl.Execute(w, material)
}

/*
	関数 CreateId()
	- 実行時のUnix時刻からミリ秒とナノ秒を切り取った6桁の数値を返す
	  サーバのメモリ制限のためユーザIDの長さは6桁とする
	  同時最大アクセス 170人まで IDが衝突とする可能性 0.03%以下
	
	引数
	- なし
	
	戻り値
	- IDとして使用する6桁の文字列(000000-999999)
*/
func CreateId() string {
	now := time.Now()
	return fmt.Sprintf("%06d", (now.UnixNano() / 1000) % 1000000)
}

/*
	関数 Check(c, err)
	- エラーが発生していたらコンソールへ出力する
	
	引数
	- c : コンソール出力用コンテキスト
	- err : error型　他の関数から渡されるエラー変数
	
	戻り値
	- なし
*/
func Check(c appengine.Context, err error) {
	if err != nil {
		c.Errorf(err.Error())
	}
}

/*
	関数 Message(gamekey, message)
	- ゲームの参加者全員にメッセージを送信する
*/
func Message(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	gamekey := r.FormValue("gamekey")
	message := r.FormValue("message")
	
	c.Debugf("GAMEKEY:%s", gamekey)
	c.Debugf("MESSAGE:%s", message)
	
	memory, err := memcache.Get(c, gamekey)
	Check(c, err)
	value := memory.Value

	players := []map[string]string{}
	msg := map[string]string{}
	json.Unmarshal(value, &players)
	json.Unmarshal([]byte(message), &msg)

	c.Debugf("FROM:%s", msg["id"])
	c.Debugf("CONTENT:%s", msg["content"])
	
	for i := range players {
		c.Debugf("target:" + players[i]["id"] + " from:" + msg["id"])
		if players[i]["id"] != msg["id"] {
			c.Debugf("SEND")
			channel.Send(c, players[i]["id"], msg["content"])
		}
	}
}