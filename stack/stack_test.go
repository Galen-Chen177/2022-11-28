package stack

import (
	"fmt"
	"testing"
)

func TestIsVailed(t *testing.T) {
	for _, v := range []string{
		"()",
		"()[]",
		"([)]",
		"((({}))",
		")()",
	} {
		fmt.Printf("%t\t%t\n", IsVailed1(v), IsVailed2(v))
	}
}
