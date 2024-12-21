package run

import (
	c "calc/pkg/decision/decode"
	"net/http"
)
func handleStart(w http.ResponseWriter, r *http.Request){
	c.Calculate(w, r)
}
func Start() {
	http.HandleFunc("/api/v1/calculate", handleStart)
	http.ListenAndServe(":8080", nil)
}