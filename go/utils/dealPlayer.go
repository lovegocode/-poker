package utils
import "fmt"
import "poker/model"
func Deal (p *model.GameData,Cards[]int){
	all:=[][]int{}
	play:=p.GetAllPlayer()
  for _,v:=range play{
	process:=append([]int(nil),Cards...)
    process=append(process,v.Hand...)
	all=append(all,process)
  }
  fmt.Println(all)
}