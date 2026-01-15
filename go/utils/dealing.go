package utils
import "poker/model"
func(T*TexasDealer) Dealing(need int,d*model.Deck)[]int{
    card:=d.AllCards[:need]
	d.Knowns=append(d.Knowns,card...)
	Allcards:=d.AllCards[need:]
	d.SetAllCards(Allcards)
	d.SendCard=append(d.SendCard,card...)
	return card
	
}