package utils

import (
	"math/rand/v2"
)
func (T*TexasDealer)ShuffleCard(){
	card:=T.Cards.AllCards
	rand.Shuffle(len(card),func(i,j int){
		card[i],card[j]=card[j],card[i]
	})
}
//抽数字索引
func RandRange(index int)int{
   randIndex:=rand.IntN(index)
   return randIndex
}