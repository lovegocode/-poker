package ranges

import (
	"poker/RangeFile"
	"sort"
)


func SortFile(hand RangeFile.AllHand)RangeFile.AllHand{
   sort.Slice(hand.HandIndex,func(i int, j int) bool{
     return hand.HandIndex[i].Win>hand.HandIndex[j].Win
   })
   return hand
}


