package service

import (
	"poker/RangeFile"
	"poker/ranges"
	"testing"
)
func TestSort(t *testing.T){
   path:="../../Range/hand.jsonl"
   dir:="../../Range"
   fileName:="sortHand.jsonl"
   data,_:=RangeFile.ReadFile(path)
   sortdata:=ranges.SortFile(data)
    for _,v:=range sortdata.HandIndex{
        hand:=v.Hand
        win:=v.Win
      ranges.SaveFile(dir,fileName,hand,win)
    }
  
}