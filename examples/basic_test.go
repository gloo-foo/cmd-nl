package nl_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-nl"
)

func ExampleNl_basic() {
	// echo "First line\nSecond line\nThird line" | nl
	output, _ := testable.Test(command.Nl(), "First line\nSecond line\nThird line\n")
	fmt.Print(output)
	// Output:
	//      1	First line
	//      2	Second line
	//      3	Third line
}
