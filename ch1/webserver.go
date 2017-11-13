package ch1

import (
	"net/http"
	"log"
	"strconv"
)

func Webserver() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 20
	color := "black"
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		switch k {
		case "cycles":
			var err error
			cycles, err = strconv.Atoi(v[0])
			if err != nil {
				cycles = 20
			}
		case "color":
			color = v[0]
		}
	}
	Lissajous(w, cycles, color)
}
