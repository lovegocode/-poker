package controller

import (
	"fmt"
	"io"
	"net/http"
)
func Message(){
	
	resp,err:=http.Get("https://svipsvip.ffzyread1.com/20240517/27297_ed48e50e/2000k/hls/mixed.m3u8")
	if err !=nil{
		fmt.Println("wrong")
	}
    defer resp.Body.Close()
    body,error:=io.ReadAll(resp.Body)
    if error!=nil{
		fmt.Println("wrong")
	}
	fmt.Println(string(body))
}