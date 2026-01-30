package Service

import (
	"math/rand/v2"
	"poker/Utils"
	"poker/model"
)

func(t*TexasDealer) Init(b*model.PlayerInfo,data*model.Board){

	
	
	t.Known() //删掉已知牌组
	t.ShuffleCard()
	t.Process()//补玩家牌
	t.Deal()

}
func (t*TexasDealer)Deal(p*model.PlayerInfo,b*model.Board){
    t.DeleteKnowDeck(p,b)
	t.Shuffle()
    
	
}
//删掉已知牌
func (t*TexasDealer)DeleteKnowDeck(p*model.PlayerInfo,b*model.Board){
	pool:=[]int{}
	public:=b.GetCards()
   pool=append(pool,public...)
   for _,v:=range p.Players{
	    hand:=v.Hand
	pool=append(pool,hand...)
   }
   for _,j:=range pool{
	for i,x:=range t.Deck.Cards{
		if j==x{
		  length:=len(t.Deck.Cards)-1
          t.Deck.Cards[i]=t.Deck.Cards[length]
		  t.Deck.Cards=t.Deck.Cards[:length]
		  break
		}

	}
   }
}
//洗牌
func (t*TexasDealer)Shuffle(){
	rand.Shuffle(len(t.Deck.Cards),func(i,j int){
	   t.Deck.Cards[i],t.Deck.Cards[j] =t.Deck.Cards[j],t.Deck.Cards[i]
	})
}
func (t*TexasDealer)PlayerCards(p*model.PlayerInfo){
	
   for i:=0;i<p.Person;i++{
     if i==p.SkipId{
		continue
	 }
	 if p.Players[i].Hand==nil{
	   hand:= Utils.RandHands(p.Players[i].PlayerRange,)
	   p.Players[i].Hand=append(p.Players[i].Hand,hand...)
	 }
   }
}
