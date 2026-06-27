// Package alias provides unprefixed names for nl command flags.
//
//	import nl "github.com/gloo-foo/cmd-nl/alias"
//	nl.Nl(nl.BodyNonEmpty, nl.Sep(": "))
package alias

import command "github.com/gloo-foo/cmd-nl"

// Nl re-exports the constructor.
var Nl = command.Nl

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

// Sep creates a -s separator flag.
var Sep = command.NlSep

// Start creates a -v starting-line-number flag.
var Start = command.NlStart

// Increment creates a -i line-number-increment flag.
var Increment = command.NlIncrement

// Width creates a -w field-width flag.
var Width = command.NlWidth

// Format creates a -n number-format flag.
var Format = command.NlFormat
