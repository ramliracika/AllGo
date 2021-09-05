package context

import (
	"context"
	"fmt"
	"testing"
)

func TestWithValue(t *testing.T) {
	ContextA := context.Background() // parent context

	ContextB := context.WithValue(ContextA, "b", "B") // child context A
	ContextC := context.WithValue(ContextA, "c", "C")

	ContextD := context.WithValue(ContextB, "d", "D") // child context B

	ContextE := context.WithValue(ContextC, "e", "E") // child context C

	ContextF := context.WithValue(ContextE, "f", "F") // child context E

	fmt.Println(ContextA)
	fmt.Println(ContextB)
	fmt.Println(ContextC)
	fmt.Println(ContextD)
	fmt.Println(ContextE)
	fmt.Println(ContextF)

	fmt.Println(ContextA.Value("a")) // get data from child to parent
	fmt.Println(ContextA.Value("f")) // nill because context A didnt have parent

	fmt.Println(ContextF.Value("c")) // data found because f child a
	fmt.Println(ContextD.Value("b"))
}
