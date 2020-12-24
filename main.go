package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main(){
	r := gee.New()
	r.Get("/",func(w http.ResponseWriter,req *http.Request){
		fmt.Fprint(w,"hello World")
	})
	r.Run(":9999")

}