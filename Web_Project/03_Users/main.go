package main

import (
	"html/template"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
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

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf(ctx, "error creating password: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	
	user := User{
		Email: req.FormValue("email"),
		Name: req.FormValue("name"),
		Username: req.FormValue("username"),
		Password: string(hashedPass),
	}
	key := datastore.NewKey(ctx, "Users", user.Username, 0, nil)
	key, err = datastore.Put(ctx, key, &user)
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}

}

func createSession(res http.ResponseWriter, req *http.Request, user User) {
	ctx := appengine.NewContext(req)
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		Path:  "/",
	}
	http.SetCookie(res, cookie)

	json, err := json.Marshal(user)
	if err != nil {
		log.Errorf(ctx, "error marshalling during user creation: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
	}
	memcache.Set(ctx, &sd)
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

}
