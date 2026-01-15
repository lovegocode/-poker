package utils
import "poker/model"
func(T*TexasDealer) Dealing(need int,d*model.Deck)[]int{
	cards:=T.cards.Get()
    card:=cards[:need]
	T.Known()=append(d.Knowns,card...)
	Allcards:=d.AllCards[need:]
	d.SetAllCards(Allcards)
	d.SendCard=append(d.SendCard,card...)
	return card
	
}