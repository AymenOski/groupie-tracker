package functions

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// it will take right path adding a security layer
	path := filepath.Clean("../" + r.URL.Path)
	fmt.Println(path)
	rr, err := os.Stat(path)
	if err != nil || rr.IsDir() {
		RenderPageNotFound(w, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, path)
}
