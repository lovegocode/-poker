package Utils

import "math/rand"
//根据范围随机抽一组牌 
func RandHands(chance int)[]int{
	result:=RangeToMap(chance)
   length:=len(result.HandIndex)
   index:=rand.Intn(length)
   cards:=result.HandIndex[index].MapHand
   return cards
}