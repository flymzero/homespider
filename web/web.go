package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/flymzero/homespider/job51"
)

var jobList = job51.GetJobDetail()

func CreateWebServer() {

	h := http.FileServer(http.Dir("/Users/program/go/src/github.com/flymzero/homespider/web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", h))

	http.HandleFunc("/index.html", homespiderHandler)
	http.HandleFunc("/hangzhou.json", jsonHandler)
	//http.HandleFunc("/jquery.js", jqueryHandler)
	http.ListenAndServe(":8000", nil)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("/Users/program/go/src/github.com/flymzero/homespider/web/hangzhou.json")
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

// func jqueryHandler(w http.ResponseWriter, r *http.Request) {
// 	b, err := ioutil.ReadFile("/Users/program/go/src/github.com/flymzero/homespider/web/jquery.js")
// 	if err != nil {
// 		fmt.Fprint(w, err)
// 	} else {
// 		fmt.Fprint(w, string(b))
// 	}
// }

func homespiderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		t, _ := template.ParseGlob("/Users/program/go/src/github.com/flymzero/homespider/web/index.html")
		t.Execute(w, map[string]interface{}{"job": jobList})
		// b, err := ioutil.ReadFile("/Users/program/go/src/github.com/flymzero/homespider/web/index.html")
		// if err != nil {
		// 	fmt.Fprint(w, err)
		// } else {
		// 	fmt.Fprint(w, string(b))
		// }
	}
}
