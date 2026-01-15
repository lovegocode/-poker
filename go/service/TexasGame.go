package service

import (
	"poker/model"
	"poker/utils"
)
type TexasGame struct{

}
func (T*TexasGame)MakeJudge()model.Judge{
 return &Texas{}
}
func (T*TexasGame)MakeDealer(b*model.Board,g*model.GameData)model.Dealer{
   return utils.NewTexasDealer(b,g)
}