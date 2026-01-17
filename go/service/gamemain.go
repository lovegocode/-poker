package service

import (
	"poker/model"
	
)
 
func GameMain( person int,result func(s int)){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池和设置自己的牌
	player:=model.NewGameData(person)
	player.Add(0,[]int{59,58})
	
    //添加公牌
	board:=model.NewBoard([]int{15,35,40})
     dealer:=texas.MakeDealer()
	 all:=dealer.Init(board,player)//接收一个要处理的二维切片
	judge:=texas.MakeJudge()
	
    judge.CallBack(result)
	judge.InitCard(all)
}