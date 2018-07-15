package handler

import (
	"net/http"
)

func ServeIndexFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "file/resource/html/index.html")
}
