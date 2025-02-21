package functions

import (
	"fmt"
	"net/http"
	"strconv"
)

func MoreInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderErrPage(w, http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "More Info for ID: %d", id)
}
