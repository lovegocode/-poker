package service

import (
	"poker/model"
	"poker/utils"
	"fmt"
)
func GameMain(){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
	//造发牌员 新牌组
	dealer:=texas.MakeDealer()
	deck,Dealer:=dealer.NewDeck()//前面deck ，后面发牌员
    //造玩家池 加玩家手牌
	player:=model.NewGameData(6)
    player.SetPlayer()//初始化id
	player.Add(0,[]int{15,30})
	
    //公牌
	border:=model.NewBoard([]int{8,9,10,25,36})
    publicCard:=border.GetBoardCards()
    // 要先处理掉已经知道的牌  
	Delete:=utils.DeletePool(publicCard,player)
	//从牌堆删除
	Dealer.Known(Delete,deck)
   fmt.Println(len(deck.AllCards))
   //发牌员洗牌 然后判断要不要发牌
   Dealer.ShuffleCard(deck)
   fmt.Println(deck.AllCards)
   //判断要不要补玩家牌和公牌
   Dealer.Process(player,deck,border)
  cards:= border.GetBoardCards()
  fmt.Println(cards)
   //创造裁判 
   judge:=texas.MakeJudge()
   //把每个玩家和公牌加一起 
   //写责任链方法
}