// Package alias provides unprefixed names for nl command flags.
//
//	import nl "github.com/gloo-foo/cmd-nl/alias"
//	nl.Nl(nl.BodyNonEmpty, nl.Sep(": "))
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-nl"
)

// Nl numbers lines from stdin; see the command package for the flag set.
func Nl(opts ...any) gloo.Command[[]byte, []byte] { return command.Nl(opts...) }

// -b flag: which lines to number.
const (
	BodyAll      = command.NlBodyAll      // number all lines (default)
	BodyNonEmpty = command.NlBodyNonEmpty // number non-empty lines only
	BodyNone     = command.NlBodyNone     // number no lines
)

// -n flag: line-number format.
const (
	FormatLN = command.NlFormatLN // left justified, no leading zeros
	FormatRN = command.NlFormatRN // right justified, no leading zeros (default)
	FormatRZ = command.NlFormatRZ // right justified, leading zeros
)

// Sep is the -s separator flag.
type Sep = command.NlSep

// Start is the -v starting-line-number flag.
type Start = command.NlStart

// Increment is the -i line-number-increment flag.
type Increment = command.NlIncrement

// Width is the -w field-width flag.
type Width = command.NlWidth

// Format is the -n number-format flag.
type Format = command.NlFormat
