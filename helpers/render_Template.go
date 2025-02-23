package helpers

import (
	"bytes"
	"net/http"

	"groupie/tools"
)

func RenderTemplates(w http.ResponseWriter, temp string, data interface{}, status int) {
	var buf bytes.Buffer
	// execute the template with buffer to check if there is an error in our templates
	err := tools.Tp.ExecuteTemplate(&buf, temp, data)
	if err != nil {
		buf.Reset()
		status = http.StatusInternalServerError
		err := tools.Tp.ExecuteTemplate(&buf, "statusPage.html", tools.ErrorInternalServerErr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(PageDeleted()))
			return
		}
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}
