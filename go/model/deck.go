package model 
type Deck struct {
	AllCards[]int
	Number[]int
	Color []int
	Knowns []int
	SendCard[]int

}
func (d*Deck)SetAllCards(allcards []int){
	d.AllCards=allcards
}
