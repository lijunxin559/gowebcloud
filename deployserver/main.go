package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
)

func reLaunch() {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>hello,I am lijunxin</h1>")
	reLaunch()
}


func main() {
	http.HandleFunc("/",firstPage)//自动调firstpage
	http.ListenAndServe(":5000",nil)
}