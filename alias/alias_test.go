package alias_test

import (
	"slices"
	"testing"

	nl "github.com/gloo-foo/cmd-nl/alias"
	"github.com/gloo-foo/testable"
)

// The alias package re-exports the nl constructor and flag symbols under
// unprefixed names. A mis-wired re-export (say, Start bound to the increment
// flag, or BodyNone bound to BodyAll) compiles cleanly, so only behavior can
// prove the wiring. Each test exercises one re-export and asserts the GNU nl
// output it must produce.

const bodyInput = "alpha\n\nbeta\n"

func run(t *testing.T, input string, opts ...any) []string {
	t.Helper()
	lines, err := testable.TestLines(nl.Nl(opts...), input)
	if err != nil {
		t.Fatal(err)
	}
	return lines
}

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_NlDefaultNumbersEveryLine(t *testing.T) {
	// Nl with no flags numbers every line, width 6, tab separator.
	assertLines(t, run(t, bodyInput),
		[]string{"     1\talpha", "     2\t", "     3\tbeta"})
}

func TestAlias_BodyAllNumbersBlankLines(t *testing.T) {
	// -b a numbers all lines, including the blank one.
	assertLines(t, run(t, bodyInput, nl.BodyAll),
		[]string{"     1\talpha", "     2\t", "     3\tbeta"})
}

func TestAlias_BodyNonEmptySkipsBlankLines(t *testing.T) {
	// -b t numbers only non-empty lines; the blank line passes through.
	assertLines(t, run(t, bodyInput, nl.BodyNonEmpty),
		[]string{"     1\talpha", "", "     2\tbeta"})
}

func TestAlias_BodyNoneNumbersNothing(t *testing.T) {
	// -b n numbers no lines: every line gets a blank width-6 field.
	assertLines(t, run(t, "alpha\nbeta\n", nl.BodyNone),
		[]string{"      \talpha", "      \tbeta"})
}

func TestAlias_SepReplacesTheSeparator(t *testing.T) {
	// -s ": " replaces the default tab separator.
	assertLines(t, run(t, "x\ny\n", nl.Sep(": ")),
		[]string{"     1: x", "     2: y"})
}

func TestAlias_StartSetsFirstNumber(t *testing.T) {
	// -v 10 starts numbering at 10.
	assertLines(t, run(t, "a\nb\n", nl.Start(10)),
		[]string{"    10\ta", "    11\tb"})
}

func TestAlias_IncrementStepsBetweenNumbers(t *testing.T) {
	// -i 5 steps the counter by 5 each line.
	assertLines(t, run(t, "a\nb\nc\n", nl.Increment(5)),
		[]string{"     1\ta", "     6\tb", "    11\tc"})
}

func TestAlias_WidthSetsFieldWidth(t *testing.T) {
	// -w 3 renders the number in a 3-wide field.
	assertLines(t, run(t, "a\nb\n", nl.Width(3)),
		[]string{"  1\ta", "  2\tb"})
}

func TestAlias_FormatLNLeftJustifies(t *testing.T) {
	// -n ln left-justifies the number within the field.
	assertLines(t, run(t, "a\n", nl.FormatLN),
		[]string{"1     \ta"})
}

func TestAlias_FormatRNRightJustifies(t *testing.T) {
	// -n rn right-justifies the number (the default form).
	assertLines(t, run(t, "a\n", nl.FormatRN),
		[]string{"     1\ta"})
}

func TestAlias_FormatRZZeroPads(t *testing.T) {
	// -n rz right-justifies with leading zeros.
	assertLines(t, run(t, "a\n", nl.FormatRZ),
		[]string{"000001\ta"})
}

func TestAlias_FormatConstructorMatchesConstant(t *testing.T) {
	// Format("rz") must wire to the same zero-padded rendering as FormatRZ.
	assertLines(t, run(t, "a\n", nl.Format("rz")),
		[]string{"000001\ta"})
}
