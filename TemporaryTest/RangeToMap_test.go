package service

import (
	"poker/Utils"
	"testing"
)
//测试过map
func TestRangeToMap(t*testing.T){
	result:=Utils.RangeToMap(1)
	for i:=0;i<len(result.HandIndex);i++{
		t.Log("\n",result.HandIndex[i].MapHand,
		result.HandIndex[i].Hand,)
	}
	t.Log("len",len(result.HandIndex))
}