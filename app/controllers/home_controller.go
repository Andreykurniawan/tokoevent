package controllers

import "net/http"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, format: "Selamat Datang di Toko Event")
}