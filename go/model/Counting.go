package model
type CountBucket struct{
	countColor[]int
	countNumber[]int
	
}
func NewCount()*CountBucket{
	c:=&CountBucket{countColor:make([]int,0),
	countNumber:make([]int,0)}
	return c
}
func(c*CountBucket) SetColor(color int){
	c.countColor=append(c.countColor,color)
}
func (c*CountBucket)GetColor()[]int{
	return c.countColor
}
func(c*CountBucket)SetNumber(number int){
	c.countNumber=append(c.countNumber,number)
}
func(c*CountBucket)GetNumber()[]int{
	return c.countNumber
}
