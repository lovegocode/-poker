package utils

import "poker/model"
func DeletePool(board[]int,g*model.GameData)[]int{
    pool:=make([]int,0)
	pool=append(pool,board...)
	for _,v:=range g.Players{
		pool=append(pool,v.Hand...)
	}
	return pool
}