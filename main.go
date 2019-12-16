package main

import (
	"github.com/bradhe/stopwatch"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
	"os"

	pd "pdviewer/pdapi"
	mt "pdviewer/maketable"
)

func main() {
	sw := stopwatch.Start()
	go func() {
		for {
			a, h := pd.GetIncidents()
			sw.Stop()
			if !h {
				sw = stopwatch.Start()
			}
			t := strings.Split(sw.String(), ".")
			mt.MakeTempl(a, t[0]+"s", pd.OnCall())
			time.Sleep(10 * time.Second)
		}
	}()
	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", pdHandler)
	log.Println("Listening...")
	http.ListenAndServe(os.Getenv("HOST"), nil)
}
func pdHandler(w http.ResponseWriter, r *http.Request) {
	post_template := template.Must(template.ParseFiles(path.Join("templates", "index.html"), path.Join("templates", "post.html")))
	if err := post_template.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
	}
}
