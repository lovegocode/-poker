package utils
import "poker/RangeFile"

type Range struct{
	Hand []Hands
}
type  Hands struct {
	Single []string
	MapHand []int
}
//玩家范围池
func RangeCards(chance int)*Range{
	 ranges:=&Range{} 
	path:="../../Range/sortvs1.jsonl"
	data,_:=RangeFile.ReadFile(path)
    index:=(len(data.HandIndex)*chance)/100
	 for i:=0;i<index;i++{
		hand:=Hands{}
		hand.Single=data.HandIndex[i].Hand
		hand.MapHand=DealMap(hand.Single)

       ranges.Hand=append(ranges.Hand,hand)
	 }
	 
	 return ranges
}