package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/web/models"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Assets")
}

func Run(rr *mux.Router) error {
	rr.Handle("/", models.Root)
	//rr.Handle("/assets/", http.StripPrefix("/server/web/templates/assets/", http.ServeFile(http.Dir("./server/web/templates/assets/"))))
	//rr.Handle("/static/", http.FileServer(http.Dir("server/web/templates/")))

	return nil
}
