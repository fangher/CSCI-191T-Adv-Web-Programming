package main

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
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
  
  
func getSession(req *http.Request) (*memcache.Item, error) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("session")
	if err != nil {
		return &memcache.Item{}, err
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}, err
	}
	return item, nil

}
