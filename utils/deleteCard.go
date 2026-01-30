package utils

func (T*TexasDealer)DeleteRange(card []int){
	d:=T.Cards
	
	for i,v:=range d.AllCards{
    for _,j:=range card{
		last:=len(d.AllCards)-1
		if v==j{
			d.AllCards[i]=d.AllCards[last]
			d.AllCards=d.AllCards[:last]
			break
		}
	}
	}
	
	
}