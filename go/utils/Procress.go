package utils

import (
	"poker/model"
	
)
func(T*TexasDealer) Process(g*model.GameData,d*model.Deck,b*model.Board){
	player:=g.GetAllPlayer()
	for _,v:=range player{
		if len(v.Hand)==0{
         cards:=T.Dealing(2,d)
		 v.Hand=append(v.Hand,cards...)
		}
	}
	card:= b.GetBoardCards()
	need:=5-len(card)
	if need!=0{
		send:=T.Dealing(need,d)
		b.Add(send)
	}
	
}