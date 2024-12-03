package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/system/time", timeHandler)

	address := "127.0.0.1:8081"
	println("Сервер запущен по адресу:", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		println("Ошибка запуска сервера:", err.Error())
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	format := r.URL.Query().Get("format")
	if format == "" {
		http.Error(w, `{"error":"format query parameter is required"}`, http.StatusBadRequest)
		return
	}

	currentTime := time.Now()

	formattedTime := currentTime.Format(format)

	response := map[string]string{"system_time": formattedTime}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, `{"error":"failed to generate JSON response"}`, http.StatusInternalServerError)
		return
	}
}
