package model 

type Dealer interface{
	  Deal()
}
type Decker interface{
	ShuffleCard()
    Dealing(need int)[]int
	Known()
	Process()   
}
