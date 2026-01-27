package service

import (
	"poker/model"
	"testing"
)
func TestMap(t *testing.T){
  model.InitMap()
  color:=57&3
  number:=57>>2
  t.Log("花色",color,"数字",number)
  hand:=model.SuitNumber(color,number)
  t.Log("手牌",hand)
}