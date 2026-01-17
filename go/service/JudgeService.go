package service

import (
	
	"poker/model"
	service "poker/service/process"
)
   
type TexasJudge struct {
  Deal [][]int
  Count *model.CountId
  DealScore model.Score
  Win int
  callback(func(score int))
}

func NewJudge()*TexasJudge{
	return &TexasJudge{}
}
func (t*TexasJudge)CallBack(f func(score int)){
    t.callback=f //回调函数
}
func(t*TexasJudge) InitCard(card[][]int){
  t.Deal=card
  t.DealCount()//计数桶
  t.DealMask()//预处理 掩码
  t.Start()
  
  t.callback(t.Win)
}
 
func (t*TexasJudge)Start(){
    var scores []int
   for i:=0;i<len(t.Count.Id);i++{
     chain:=service.NewChain()
     d:=t.Count.Id[i]
     _,score:=chain.Process(d)
     scores=append(scores,score)
   }
  t.Win+= t.Compare(scores)
  
}
func (t*TexasJudge)Compare(scores[]int)int{
    
   my:=scores[0]
   win:=0
   for i:=1;i<len(scores);i++{
    if my<scores[i]{
      break
    } else{
      win++
    }
   }
   if win==len(scores)-1{
    return 1
   }
   return 0
}
