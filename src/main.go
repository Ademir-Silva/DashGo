//go:build ignore

package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "DashGo - CyberMakers Network!!")
	})

	http.ListenAndServe(":8080", nil)
}