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

// settings is the fully-resolved configuration for one Nl run: every optional
// flag has been collapsed to a concrete value.
type settings struct {
	body      nlBodyFlag
	sep       string
	format    nlFormatFlag
	start     int
	increment int
	width     int
}

// resolve collapses the raw, partially-defaulted flags into concrete settings,
// applying each GNU nl default where the caller gave no value.
func resolve(f flags) settings {
	return settings{
		body:      orBody(f.body, NlBodyAll),
		sep:       orString(f.sep, defaultSep),
		start:     orInt(f.start, defaultStart),
		increment: orInt(f.increment, defaultIncrement),
		width:     orInt(f.width, defaultWidth),
		format:    orFormat(f.format, NlFormatRN),
	}
}

// orBody returns v when set, else the fallback.
func orBody(v, fallback nlBodyFlag) nlBodyFlag {
	if v == "" {
		return fallback
	}
	return v
}

// orFormat returns v when set, else the fallback.
func orFormat(v, fallback nlFormatFlag) nlFormatFlag {
	if v == "" {
		return fallback
	}
	return v
}

// orString returns v when set, else the fallback.
func orString(v, fallback string) string {
	if v == "" {
		return fallback
	}
	return v
}

// orInt dereferences v when set, else returns the fallback.
func orInt(v *int, fallback int) int {
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
	s := resolve(gloo.NewParameters[gloo.File, flags](opts...).Flags)
	return patterns.StatefulMap(func() func([]byte) ([]byte, error) {
		// Pre-subtract so the first increment lands exactly on start.
		n := s.start - s.increment
		return func(line []byte) ([]byte, error) {
			return s.number(&n, line), nil
		}
	})
}

// number renders one line. n holds the running counter (advanced in place for
// each numbered line). Blank lines under body "t" pass through untouched.
func (s settings) number(n *int, line []byte) []byte {
	if s.body == NlBodyNone {
		return s.compose(blank(s.width), line)
	}
	if s.body == NlBodyNonEmpty && len(line) == 0 {
		return line
	}
	*n += s.increment
	return s.compose(formatNumber(*n, s.width, s.format), line)
}

// compose joins a rendered number field and the line with the separator.
func (s settings) compose(field string, line []byte) []byte {
	return fmt.Appendf(nil, "%s%s%s", field, s.sep, string(line))
}

// blank renders a width-wide run of spaces, used for unnumbered lines (body "n").
func blank(width int) string {
	return fmt.Sprintf("%*s", width, "")
}

// formatNumber formats a line number according to the given format and width.
func formatNumber(n, width int, format nlFormatFlag) string {
	switch format {
	case NlFormatLN:
		return fmt.Sprintf("%-*d", width, n)
	case NlFormatRZ:
		return fmt.Sprintf("%0*d", width, n)
	default: // NlFormatRN
		return fmt.Sprintf("%*d", width, n)
	}
}
