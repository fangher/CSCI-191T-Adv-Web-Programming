package main

import (
	"fmt"
	"net/http"
	"strings"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", handleStuff)
}

type Stuff struct {
	Something string
	URL string
}

func handleStuff(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		something := strings.Split(req.URL.Path, "/")[1]
		showStuff(res, req, something)
		return
	}

	if req.Method == "POST" {
		saveStuff(res, req)
		return
	}

	listStuffs(res, req)
}

func listStuffs(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Stuff").Order("Something")

	html := "<h2>Database</h2>"

	iterator := q.Run(ctx)
	for {
		var entity Stuff
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		html += `
			<h3>` + entity.Something + `</h3>
			<br/>
			<img src='` + entity.URL + `' />
			<br/>
		`
	}

func showStuff(res http.ResponseWriter, req *http.Request, something string) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Stuff", something, 0, nil)
	var entity Stuff
	err := datastore.Get(ctx, key, &entity)
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<h2>` + entity.Something + `</h2>
		<br/>
		<img src='` + entity.URL + `' />
		<br/>
	`)
}

func saveStuff(res http.ResponseWriter, req *http.Request) {
	something := req.FormValue("something")
	url := req.FormValue("url")
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Stuff", something, 0, nil)
	entity := Stuff{
		Something: something,
		URL: url,
	}

	_, err := datastore.Put(ctx, key, &entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	http.Redirect(res, req, "/", 302)
}
