package utils

import "poker/model"

func DealCount(cards[]int){
   c:=model.NewCount()

   for _,v:=range cards{
      c.SetColor(v&3)
      c.SetNumber(v>>2)
   }
}