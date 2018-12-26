package main

import (
	"fmt"
	"shuffler"
)

type intSlice []int

func (is intSlice) Len() int  {
	return len(is)
}

func (is intSlice) Swap(i,j int)  {
	/*	Swap code old style
		temp := is[i]
		is[i] = is[j]
		is[j] = temp*/

	//concurrent swapping in go
	is[i] , is[j] = is[j], is[i]
}


type weightedString struct {
	weight int
	s string
}

type stringSlice []weightedString

func (ss stringSlice) Len() int  {
	return len(ss)
}

func (ss stringSlice) Swap(i,j int)  {
	ss[i] , ss[j] = ss[j], ss[i]
}

func (ss stringSlice ) Weight(i int) int {
	return ss[i].weight
}



func main()  {

	is := intSlice{1,2,3,4,5,6}
	shuffler.Shuffle(is)
	fmt.Printf("%q\n\n", is)
	fmt.Printf("%v\n\n", is)
	fmt.Printf("Looped %d time\n\n", shuffler.GetCount())

	fmt.Printf("====================\n")
	fmt.Printf("Weighted Shuffle")
	i := stringSlice{{100, "hello"},{200, "world!"},
		{10,"Goodbye!"}}

	shuffler.WeightedShuffle(i)
	fmt.Printf("%v\n\n", i)

}
