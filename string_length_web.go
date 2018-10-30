package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", StringPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func StringPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, err := template.ParseFiles("string.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	} else {
		r.ParseForm()
		var output string
		var name string
		name = fmt.Sprintf(r.Form.Get("myString"))
		var l int
		l = len(name)
		output = fmt.Sprintf("%s has %d characters", name, l)
		fmt.Fprint(w, output)
	}
}
