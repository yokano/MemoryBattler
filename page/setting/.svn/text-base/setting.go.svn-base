package setting

import(
	"net/http"
	"appengine"
	"appengine/user"
	"fmt"
	. "library"
)


func Setting(w http.ResponseWriter, r *http.Request) {
	var context appengine.Context = appengine.NewContext(r)
	var player *user.User = user.Current(context)
		
	type LoginTerm struct {
		Message string;
		Theme string
	}
	material := make(map[string]string)
	material["Theme"] = "d"
	
	url,_ := user.LoginURL(context, "/setting")
	if player == nil {
		material["Message"] = fmt.Sprintf("<a href=\"%s\">Login</a>", url)
	} else {
		material["Message"] = fmt.Sprintf("<a href=\"%s\">Logout</a>", url)
	}
	Output(w, "page/setting/setting.html", material)
}