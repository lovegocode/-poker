package service

import (
	"poker/model"
	"fmt"
)
 
func GameMain(data*model.Begin,result func(s float32)){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池和设置自己的牌
	player:=model.NewGameData(data.Person)
	player.Add(data.Id,data.Hand,1)//1先代表全下
	 id:=data.Id//知道自己哪个位置
	player.SetAction(3,1)
	player.SetAction(5,1)
     
    //添加公牌
	board:=model.NewBoard(data.PublicCard,data.Person)
     dealer:=texas.MakeDealer()
	
	 dealer.Init(board,player)
	judge:=texas.MakeJudge()
	fmt.Println("这里没问题",len(board.Entry))
    judge.CallBack(result)
	judge.InitCard(id,board)
}