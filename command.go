package command

import (
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// Default values for the optional numbering flags, applied when the caller
// supplies no override.
const (
	defaultSep       = "\t"
	defaultStart     = 1
	defaultIncrement = 1
	defaultWidth     = 6
)

// lineNumber is a value on the line-number axis: the running counter that the
// -v start seeds and the -i increment advances.
type lineNumber int

// numberWidth is the character width of the rendered line-number field (-w).
type numberWidth int

// settings is the fully-resolved configuration for one Nl run: every optional
// flag has been collapsed to a concrete value.
type settings struct {
	sep       string
	body      NlBody
	format    NlFormat
	start     lineNumber
	increment lineNumber
	width     numberWidth
}

// resolve collapses the raw, partially-defaulted flags into concrete settings,
// applying each GNU nl default where the caller gave no value.
func resolve(f flags) settings {
	return settings{
		body:      orZero(f.body, NlBodyAll),
		sep:       string(orZero(f.sep, defaultSep)),
		start:     lineNumber(orDefault(f.start, defaultStart)),
		increment: lineNumber(orDefault(f.increment, defaultIncrement)),
		width:     numberWidth(orDefault(f.width, defaultWidth)),
		format:    orZero(f.format, NlFormatRN),
	}
}

// orZero returns v when set, else the fallback. The type's zero value marks
// "not set" for flags whose zero is never a meaningful choice.
func orZero[T comparable](v, fallback T) T {
	var zero T
	if v == zero {
		return fallback
	}
	return v
}

// orDefault dereferences v when set, else returns the fallback. The numeric
// flags mark "not set" with a nil pointer so an explicit zero stays a legal
// value.
func orDefault[T any](v *T, fallback T) T {
	if v == nil {
		return fallback
	}
	return *v
}

// Nl numbers lines from stdin.
//
// Flags:
//   - NlBodyAll, NlBodyNonEmpty, NlBodyNone (-b): which lines to number (default: all)
//   - NlSep(s) (-s): separator between number and line (default: "\t")
//   - NlStart(n) (-v): starting line number (default: 1)
//   - NlIncrement(n) (-i): line number increment (default: 1)
//   - NlWidth(n) (-w): field width for line numbers (default: 6)
//   - NlFormat(f) / NlFormatLN / NlFormatRN / NlFormatRZ (-n): number format (default: rn)
func Nl(opts ...any) gloo.Command[[]byte, []byte] {
	s := resolve(fold(opts))
	return patterns.StatefulMap(func() func([]byte) ([]byte, error) {
		// Pre-subtract so the first increment lands exactly on start.
		n := s.start - s.increment
		return func(line []byte) ([]byte, error) {
			out, next := s.number(n, line)
			n = next
			return out, nil
		}
	})
}

// number renders one line against the running counter n, returning the
// rendered line and the counter to carry to the next line (advanced only when
// the line was numbered). Blank lines under body "t" pass through untouched.
func (s settings) number(n lineNumber, line []byte) ([]byte, lineNumber) {
	if s.body == NlBodyNone {
		return s.compose(blank(s.width), line), n
	}
	if s.body == NlBodyNonEmpty && len(line) == 0 {
		return line, n
	}
	n += s.increment
	return s.compose(formatNumber(n, s.width, s.format), line), n
}

// compose joins a rendered number field and the line with the separator.
func (s settings) compose(field string, line []byte) []byte {
	return fmt.Appendf(nil, "%s%s%s", field, s.sep, string(line))
}

// blank renders a width-wide run of spaces, used for unnumbered lines (body "n").
func blank(width numberWidth) string {
	return fmt.Sprintf("%*s", int(width), "")
}

// formatNumber formats a line number according to the given format and width.
func formatNumber(n lineNumber, width numberWidth, format NlFormat) string {
	switch format {
	case NlFormatLN:
		return fmt.Sprintf("%-*d", int(width), int(n))
	case NlFormatRZ:
		return fmt.Sprintf("%0*d", int(width), int(n))
	default: // NlFormatRN
		return fmt.Sprintf("%*d", int(width), int(n))
	}
}
