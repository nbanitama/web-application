package core

import (
	"net/http"
)

func (m *Module) ServeIndexFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "file/resource/html/index.html")

}
