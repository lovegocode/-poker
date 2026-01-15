package model 
type Dealer interface{
   Init()
}
type Decker interface{

	ShuffleCard()
    Dealing(need int)[]int
	Known(pool[]int)
	Process(g*GameData,d*Deck,b*Board)
    
}
