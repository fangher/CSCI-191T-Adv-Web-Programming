package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var templ *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/form/login", Login)
	r.GET("/form/signup", Signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	r.POST("/api/login", loginProcess)
	r.GET("/api/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	templ = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "home.html")
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "login.html")
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "signup.html")
}
