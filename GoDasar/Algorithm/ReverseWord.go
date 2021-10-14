package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) {

	a := strings.Split(str, " ")
	var temp []string
	for i := len(a) - 1; i >= 0; i-- {
		temp = append(temp, a[i])
	}
	res := strings.Join(temp," ")
	fmt.Print(res)
}
func main()  {

	ReverseWords("hello world!")

}


//Expect(ReverseWords("hello world!")).To(Equal("world! hello"))
//Expect(ReverseWords("yoda doesn't speak like this")).To(Equal("this like speak doesn't yoda"))
//Expect(ReverseWords("foobar")).To(Equal("foobar"))
//Expect(ReverseWords("kata editor")).To(Equal("editor kata"))
//Expect(ReverseWords("row row row your boat")).To(Equal("boat your row row row"))
