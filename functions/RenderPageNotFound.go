package functions

import (
	"fmt"
	"net/http"

	"Student/tools"
)

func RenderPageNotFound(w http.ResponseWriter, Status int) {
	w.WriteHeader(Status)
	type msgStatus struct {
		Messages map[int]string
	}

	statusMessages := msgStatus{
		Messages: map[int]string{
			400: "Bad request.",
			404: "We can't find the page you're looking for.",
			405: "Method not allowed. Please check the request method.",
			500: "Something went wrong on our side. Please try again later.",
		},
	}
	data := map[string]interface{}{
		"Status":        Status,
		"RollBackerror": statusMessages.Messages[Status],
	}
	err := tools.Temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
		return
	}
}
