package nl_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-nl"
	"github.com/gloo-foo/testable"
)

func ExampleNl_fromFile_basic() {
	// nl testdata/text.txt
	output, _ := testable.Test(command.Nl(), "First line\nSecond line\nThird line\n")
	fmt.Print(output)
	// Output:
	//      1	First line
	//      2	Second line
	//      3	Third line
}
