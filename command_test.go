package command_test

import (
	"testing"

	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/assertion"

	command "github.com/gloo-foo/cmd-nl"
)

func TestNl_Basic(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(), "hello\nworld\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"     1\thello",
		"     2\tworld",
	})
}

func TestNl_MultipleLines(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(), "a\nb\nc\nd\ne\n")
	assertion.NoError(t, err)
	assertion.Count(t, lines, 5)
	assertion.Contains(t, lines, "     1\ta")
	assertion.Contains(t, lines, "     5\te")
}

func TestNl_BodyNonEmpty(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlBodyNonEmpty), "a\n\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"     1\ta",
		"",
		"     2\tb",
	})
}

func TestNl_BodyNone(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlBodyNone), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Count(t, lines, 2)
	assertion.NotContains(t, lines, "1")
}

func TestNl_CustomSeparator(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlSep(": ")), "x\ny\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"     1: x",
		"     2: y",
	})
}

func TestNl_UnknownOptionsIgnored(t *testing.T) {
	// nl takes no positional arguments: values of any non-flag type are
	// ignored and the GNU defaults still apply.
	lines, err := testable.TestLines(command.Nl("ignored", 42, nil), "hello\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"     1\thello",
	})
}

func TestNl_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(), "")
	assertion.NoError(t, err)
	assertion.Empty(t, lines)
}

func TestNl_BodyNonEmpty_AllEmpty(t *testing.T) {
	output, err := testable.Test(command.Nl(command.NlBodyNonEmpty), "\n\n\n")
	assertion.NoError(t, err)
	assertion.True(t, output == "\n\n\n", "expected three empty lines")
}

func TestNl_BodyNonEmpty_Mixed(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlBodyNonEmpty), "\nfirst\n\nsecond\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"",
		"     1\tfirst",
		"",
		"     2\tsecond",
	})
}

func TestNl_StartFlag(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlStart(10)), "a\nb\nc\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"    10\ta",
		"    11\tb",
		"    12\tc",
	})
}

func TestNl_IncrementFlag(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlIncrement(5)), "a\nb\nc\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"     1\ta",
		"     6\tb",
		"    11\tc",
	})
}

func TestNl_WidthFlag(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlWidth(3)), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"  1\ta",
		"  2\tb",
	})
}

func TestNl_WidthWide(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlWidth(10)), "x\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"         1\tx",
	})
}

func TestNl_FormatLN(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlFormat("ln")), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"1     \ta",
		"2     \tb",
	})
}

func TestNl_FormatRN(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlFormat("rn")), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"     1\ta",
		"     2\tb",
	})
}

func TestNl_FormatRZ(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlFormat("rz")), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"000001\ta",
		"000002\tb",
	})
}

func TestNl_StartAndIncrement(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlStart(100), command.NlIncrement(10)), "a\nb\nc\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"   100\ta",
		"   110\tb",
		"   120\tc",
	})
}

func TestNl_FormatRZWithWidth(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlFormat("rz"), command.NlWidth(4)), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"0001\ta",
		"0002\tb",
	})
}

func TestNl_BodyNoneWithWidth(t *testing.T) {
	lines, err := testable.TestLines(command.Nl(command.NlBodyNone, command.NlWidth(3)), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{
		"   \ta",
		"   \tb",
	})
}

func TestNl_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		opts     []any
		input    string
		expected []string
	}{
		{
			name:  "single line",
			input: "only\n",
			expected: []string{
				"     1\tonly",
			},
		},
		{
			name:  "body all with empties",
			opts:  []any{command.NlBodyAll},
			input: "a\n\nb\n",
			expected: []string{
				"     1\ta",
				"     2\t",
				"     3\tb",
			},
		},
		{
			name:  "custom sep with body non-empty",
			opts:  []any{command.NlBodyNonEmpty, command.NlSep(" | ")},
			input: "x\n\ny\n",
			expected: []string{
				"     1 | x",
				"",
				"     2 | y",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := testable.TestLines(command.Nl(tt.opts...), tt.input)
			assertion.NoError(t, err)
			assertion.Lines(t, lines, tt.expected)
		})
	}
}
