package command

// NlBody selects which lines are numbered.
// Maps to the -b flag: "a" (all, default), "t" (non-empty only), "n" (none).
type NlBody string

const (
	NlBodyAll      NlBody = "a"
	NlBodyNonEmpty NlBody = "t"
	NlBodyNone     NlBody = "n"
)

// NlSep sets the separator between the line number and the line text.
// Maps to the -s flag. Default is "\t".
type NlSep string

// NlStart sets the starting line number (-v flag, default 1).
type NlStart int

// NlIncrement sets the line-number increment (-i flag, default 1).
type NlIncrement int

// NlWidth sets the field width for line numbers (-w flag, default 6).
type NlWidth int

// NlFormat sets the line-number format (-n flag).
// Valid values: "ln" (left justified), "rn" (right justified, default), "rz" (right justified zero-padded).
type NlFormat string

const (
	NlFormatLN NlFormat = "ln"
	NlFormatRN NlFormat = "rn"
	NlFormatRZ NlFormat = "rz"
)

// flags is the raw, partially-defaulted option set folded from an Nl call's
// option values. The numeric fields are pointers so an explicit zero remains
// distinguishable from "not set".
type flags struct {
	start     *int
	increment *int
	width     *int
	body      NlBody
	sep       NlSep
	format    NlFormat
}

// with folds one option value into the flag set. Values of any other type are
// ignored: nl takes no positional arguments.
func (f flags) with(o any) flags {
	switch v := o.(type) {
	case NlBody:
		f.body = v
	case NlSep:
		f.sep = v
	case NlStart:
		n := int(v)
		f.start = &n
	case NlIncrement:
		n := int(v)
		f.increment = &n
	case NlWidth:
		n := int(v)
		f.width = &n
	case NlFormat:
		f.format = v
	}
	return f
}

// fold collapses the Nl option values into the raw flag set.
func fold(opts []any) flags {
	var f flags
	for _, o := range opts {
		f = f.with(o)
	}
	return f
}
