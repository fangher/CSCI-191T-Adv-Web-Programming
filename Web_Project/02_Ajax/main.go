package main

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine"
)

var templ *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/login", login)
	r.GET("/signup", signup)
	r.GET("/view", view)
	r.GET("/write", write)
	r.GET("/profile", profile)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	templ = template.New("roottemplate")
	templ = template.Must(templ.ParseGlob("templates/*.html"))
  }

func login(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(res, "login.html", nil)
  }

func signup(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(res, "signup.html", nil)
  }

func view(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(res, "view.html", nil)
  }

func write(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(res, "write.html", nil)
  }

func profile(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(res, "profile.html", nil)
  }
  
func profilename(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	sd := sessionInfo(req)
	var user User
	user.Username = ps.ByName("name")
	if user.Username != sd.Username {
		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Users", user.Username, 0, nil)
		err := datastore.Get(ctx, key, &user)
		if err != nil {
			panic(err)
		}
	} else {
		user = sd.User
	}

}

func checkName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	log.Infof(ctx, "REQUEST BODY: %v", sbs)
	var user User
	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
	err = datastore.Get(ctx, key, &user)
	log.Infof(ctx, "ERR: %v", err)
	if err != nil {
		fmt.Fprint(res, "false")
		return
	} else {
		fmt.Fprint(res, "true")
	}
}

func checkEmail(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	q := datastore.NewQuery("Users").Filter("Email =", sbs)
	var u []User
	_, err = q.GetAll(ctx, &u)
	if err != nil {
		panic(err)
	}
	if len(u) > 0 {
		fmt.Fprint(res, "true")
	} else {
		fmt.Fprint(res, "false")
		return
	}
}


}
