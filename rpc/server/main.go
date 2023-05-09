package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		fmt.Println("path: ", request.URL.Path)
		a, _ := strconv.Atoi(request.Form["a"][0])
		b, _ := strconv.Atoi(request.Form["b"][0])
		writer.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		writer.Write(jData)
	})

	_ = http.ListenAndServe(":8080", nil)
}
