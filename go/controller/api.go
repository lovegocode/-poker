package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"poker/model/httpModel"
) 
func Route() {
	http.HandleFunc("/cal",DealRequest)
	http.HandleFunc("/go",Go)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
func DealRequest(w http.ResponseWriter, req*http.Request){
	fmt.Println("收到请求")
	data:=httpModel.NewData()
	err:=json.NewDecoder(req.Body).Decode(data)
	if err!=nil{
		return 
	}
	fmt.Println(data)
	 reqs:=Begining(data)
	 fmt.Println(reqs)
    w.WriteHeader(http.StatusOK)
    
}
func Go(w http.ResponseWriter,req*http.Request){
	fmt.Println("第二个请求收到")
	io.WriteString(w,"i love go ")
}