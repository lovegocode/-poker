package model
type Judge interface{
	DealCount()
	InitCard(card[][]int)
}
type Score interface{
	Processing(d*Detail)(bool,int)
}