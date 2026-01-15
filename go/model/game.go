package model
type Game interface{
	MakeJudge()Judge
	MakeDealer()Dealer
}
type GameData struct{
   Players []*Player 
}
func NewGameData(person int)*GameData{
	G:=&GameData{Players:make([]*Player,0,person)}
	return G
}
func (g*GameData)Set(p []*Player){
	g.Players=p
}
func (g*GameData)Add(id int,hand[]int)*GameData{
      p:=&Player{Id:id,Hand:hand}
	g.Players=append(g.Players,p)
	return g
}
func (g*GameData)GetAllPlayer()[]*Player{
   return g.Players
}
func (g*GameData)IndexPlayer(index int)*Player{
	   return g.Players[index]
}