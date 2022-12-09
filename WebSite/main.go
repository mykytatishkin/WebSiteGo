// 	http://localhost:8080/
//
//	WebSiteGo
// 	main.go
//
//	Created by Mykyta Tishkin on 09.12.2022

package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

func main() {

	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("Hello, World!")
}
