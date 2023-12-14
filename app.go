package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("🚀程式啟動，滑鼠點擊右邊網址開啟網頁 http://127.0.0.1:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//解決 CORS 問題
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "hx-current-url,hx-request")

		if r.Method == "GET" {
			fmt.Fprintf(w, "<p>🚀GET</p>")
		}
		if r.Method == "POST" {
			fmt.Fprintf(w, "<p>🚀POST</p>")
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
