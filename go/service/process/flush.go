package service

import (
	"math/bits"
	"poker/model"
)

type Flush struct {
}

func (f *Flush) Processing(d *model.Detail)(bool,int){
	for i:=0;i<4;i++{
	mask:=d.MaskColor[i]
	if bits.OnesCount16(mask)>=5{
       
	}
	}
	return false,0
}