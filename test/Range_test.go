package service

import (
	"poker/utils"
	"testing"
)
func TestRange(t *testing.T){
  ranges:= utils.RangeCards(40)//概率 40代表40%
   
   index:=utils.RandRange(len(ranges.Hand))
   hand:=ranges.Hand[index].MapHand
    deal:=utils.NewTexasDealer()
    deck:=utils.NewDeck()
    deal.Cards=deck
    t.Log(hand)
    deal.DeleteRange(hand)

   t.Log(deal.Cards.AllCards)
   t.Log(len(deal.Cards.AllCards))
}