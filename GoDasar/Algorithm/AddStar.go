package main

import (
	"fmt"
	"sort"
	"strings"

)

func TwoSort(arr []string){
 sort.Slice(arr, func(i, j int) bool {
	 return arr[i] < arr[j]
 })

temp := strings.Split(arr[0],"")

res := strings.Join(temp,"***")

fmt.Println(res)

}

func  main()  {
	var s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
	TwoSort(s)

}

//s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
//Expect(TwoSort(s)).To(Equal("b***i***t***c***o***i***n"))
//s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
//Expect(TwoSort(s)).To(Equal("a***r***e"))
//s = []string{"lets", "talk", "about", "javascript", "the", "best", "language"}
//Expect(TwoSort(s)).To(Equal("a***b***o***u***t"))
//s = []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}
//Expect(TwoSort(s)).To(Equal("c***o***d***e"))
//s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
//Expect(TwoSort(s)).To(Equal("L***e***t***s"))