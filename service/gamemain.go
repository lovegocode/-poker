package service

import (
	"poker/model"
	
)
 //反回judge计算的结果集
func GameMain(data*model.Begin)model.Result{
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池和设置自己的牌
	player:=model.NewGameData(data.Person)
	player.Add(data.Id,data.Hand,1)//1先代表全下
	 id:=data.Id//知道自己哪个位置
	SetPlayer(player,data.Action)
	
    //添加公牌
	board:=model.NewBoard(data.PublicCard,data.Person)
     dealer:=texas.MakeDealer()
	
	 dealer.Init(board,player)
	judge:=texas.MakeJudge()
	
    
     result:=judge.InitCard(id,board)
	 return result
}