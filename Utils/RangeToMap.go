package Utils
//处理好手牌map后的
func RangeToMap(chance int)AllHand{
   
	hand:=PlayerRange(chance)
	for i:=0;i<len(hand.HandIndex);i++{
		pool:=hand.HandIndex[i].Hand
		result:=DealMap(pool)
		hand.HandIndex[i].MapHand=result
	}
   return hand
}