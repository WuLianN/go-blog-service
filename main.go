package main

import (
	"net/http"
	"time"
	"github.com/WuLianN/go-blog/routers"
)

func main() {
	routersInit := routers.SetupRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}