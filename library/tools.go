/*
	tools.go 2012/6/29
	yuta.okano@gmail.com
	
	汎用的に使用する関数群
*/

package library

import (
	"text/template"
	"net/http"
	"time"
	"fmt"
)

/*
	関数 Output()
	- HTMLテンプレートにデータを置換してクライアントへ出力する

	引数
	- path : HTMLテンプレートのファイル名(page/*** の *** 部分)
	- material : 置換するデータを含むマップ
	
	戻り値
	- なし : クライアントの画面にHTMLを出力
*/
func Output(w http.ResponseWriter, path string, material map[string]string) {
	tmpl,err := template.ParseFiles(path)
	if err != nil {
		err.Error()
	}
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
