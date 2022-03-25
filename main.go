package main

import (
	"fmt"
	"job/err_no_rows/api/album"
	"job/err_no_rows/dal"
	"log"
	"net/http"
)

type handerFunc func(http.ResponseWriter, *http.Request)

func (f handerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你好"))
}

func main() {
	err := dal.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dal.Close()
	http.Handle("/", handerFunc(hello))
	http.Handle("/album", handerFunc(album.Album))
	err = http.ListenAndServe("localhost:8090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
