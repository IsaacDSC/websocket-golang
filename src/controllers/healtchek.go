package controllers

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HealthCheck")
	w.Write([]byte("OK"))
}
