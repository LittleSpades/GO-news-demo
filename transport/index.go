package transport

import (
	"bytes"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	tpl := template.Must(template.ParseFiles("index.html"))
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	buf.WriteTo(w)
}
