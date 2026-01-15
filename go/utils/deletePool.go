package utils

func(T*TexasDealer) DeletePool()[]int{
    pool:=make([]int,0)
	cards:=T.board.GetBoardCards()
	player:=T.data.GetAllPlayer()
	pool=append(pool,cards...)
	for _,v:=range player{
		pool=append(pool,v.Hand...)
	}
	return pool
}