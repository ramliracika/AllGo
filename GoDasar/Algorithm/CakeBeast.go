package main

import (
	"fmt"
	"strings"
)

func Feast(beast string, dish string) bool  {
	// your code here

	b := strings.Split(beast,"")
	d := strings.Split(dish,"")

	if b[0] == d[0] && b[len(b)-1] == d[len(d)-1] {
		return true
	}
	return false

}

func main()  {
	fmt.Println(Feast("brown bear", "bear claw"))
}

//Expect(Feast("great blue heron", "garlic naan")).To(BeTrue(), "Testing 'great blue heron' vs 'garlic naan'")
//Expect(Feast("chickadee", "chocolate cake")).To(BeTrue(), "Testing 'chickadee' vs 'chocolate cake'")
//Expect(Feast("brown bear", "bear claw")).To(BeFalse(), "Testing 'brown bear' vs 'bear claw'")

