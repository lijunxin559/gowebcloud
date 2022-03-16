package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>hello,I am lijunxin</h1>")
}


func main() {
	http.HandleFunc("/",firstPage)//自动调firstpage
	http.ListenAndServe(":8000",nil)
}

