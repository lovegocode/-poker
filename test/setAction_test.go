package service

import (
	"poker/model"
	"poker/service"
	"testing"
)

func TestAction(t *testing.T){
	
	 player:=model.NewGameData(6)
    action:=[]int{1,3,5}
     service.SetPlayer(player,action)
	t.Log(player.Players[4])
}