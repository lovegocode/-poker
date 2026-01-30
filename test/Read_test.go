package service

import (
	"poker/RangeFile"
	
	"testing"
)

func TestReadFile(t *testing.T){
   path:="../../Range/Hands.jsonl"
	data,err:=RangeFile.ReadFile(path)
  if err!=nil{
    t.Log("错误",err)
  }
 for _,v:=range data.HandIndex{
  t.Log("datawin",v)
 }
  
}