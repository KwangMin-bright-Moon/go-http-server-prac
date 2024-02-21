package main

import (
	"net/http"

	"github.com/KwangMin-bright-Moon/tucker/WEB1/myapp"
)



func main(){
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
} 