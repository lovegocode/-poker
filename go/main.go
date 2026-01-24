package main

import (
	"poker/controller"
	"poker/utils"
)

func main(){
    controller.Route()
	pool:=string[]{"A8",}
	utils.DealMap(pool)
} 