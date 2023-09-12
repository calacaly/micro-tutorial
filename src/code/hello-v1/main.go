package main

import (
	"fmt"
	"log"
	"net/http"
)

var routers []string

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("getting request for host: ", r.Host)
	fmt.Fprintf(w, "hello world")
}

func root(w http.ResponseWriter, r *http.Request) {
	ret_tpl := "<html><body>"
	for _, router := range routers {
		ret_tpl += fmt.Sprintf(
			"<a href=%s>%s<br>",
			router,
			router,
		)
	}

	ret_tpl += "</body></html>"
	fmt.Fprintf(w, "%s", ret_tpl)
}

func main() {
	log.Default()
	log.Println("starting")
	http.HandleFunc("/", root)
	http.HandleFunc("/v1/hello", hello)

	routers = append(routers, "/")
	routers = append(routers, "/v1/hello")

	http.ListenAndServe(":9090", nil)
}
