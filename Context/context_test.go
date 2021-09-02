package context

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(t *testing.T) {

	background := context.Background()
	fmt.Println(background)

	// context background is never cancelled,timeout, no value
	// use in the beginning request

	todo := context.TODO()
	fmt.Println(todo)

	// context todo same like background but not clear what context to use
}
