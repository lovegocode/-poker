package service

import "poker/model"

type TexasJudge struct {
  Deal [][]int
  Count model.CountId
  
}

func NewJudge()*TexasJudge{
	return &TexasJudge{}
}
func(t*TexasJudge) InitCard(card[][]int){
  t.Deal=card
  t.DealCount()//计数桶
  t.DealMask()//预处理 掩码
  
}
