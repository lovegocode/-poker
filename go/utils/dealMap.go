package utils

import "poker/model"
import "fmt"
func DealMap(pool[]string){
	_=make([]int,0)
	for i:=0;i<len(pool);i++{
	    deal:=pool[i][:1] 
        color:=model.SuitMap[deal]

		fmt.Println(color)
	}
}