package ranges

import (
	"fmt"
	"poker/model"
	"poker/utils"
)


func Combination (){
	deck:=utils.NewDeck()
	cards:=deck.AllCards
     model.InitMap()
	 z:=0
	 
	for i:=len(cards)-1;i>=1;i--{
		for j:=i-1;j>=0;j--{
           Suit1:=cards[i]&3
		   Suit2:=cards[j]&3
		   Number1:=cards[i]>>2
		   Number2:=cards[j]>>2 
		 hand1:= model.SuitNumber(Suit1,Number1)
         hand2:=model.SuitNumber(Suit2,Number2)
		 hand:=[]string{}
		hand=append(hand, hand1)
		hand=append(hand,hand2) 
         
		 model.SetHandMap(z,hand,0)
		 fmt.Println(model.HandMap[z])
	     z++
		}
	}

}