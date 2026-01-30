package service

import (
	"poker/model"
	"poker/model/httpModel"
	
	"poker/utils"
)

//业务入口 蒙特卡洛计算德州
func CalRate(cal*httpModel.CalData)model.Result{
   result:=model.Result{}
   data:= ToDealData(cal)//处理好的对象
    _=GameMain(data)//每一次的

   return result
}
//把数据简单处理了 该map的map
func ToDealData(cal*httpModel.CalData)*model.Begin{
	hand:=cal.Cards.Hand
	public:=cal.Cards.PublicCards
	Handmap:=utils.DealMap(hand)
	PublicMap:=utils.DealMap(public)
	begin:=&model.Begin{
	  Person: cal.Table.Person,
	  Id: cal.Cards.Position,
      Hand: Handmap,
      PublicCard: PublicMap,
	  Frequency: 10000,
	  Action:cal.Table.PlayerAction,
	}
	return begin
}