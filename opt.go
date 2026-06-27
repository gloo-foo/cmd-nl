package command

// nlBodyFlag controls which lines are numbered.
// Maps to the -b flag: "a" (all, default), "t" (non-empty only), "n" (none).
type nlBodyFlag string

const (
	NlBodyAll      nlBodyFlag = "a"
	NlBodyNonEmpty nlBodyFlag = "t"
	NlBodyNone     nlBodyFlag = "n"
)

func (f nlBodyFlag) Configure(flags *flags) { flags.body = f }

// nlSepFlag controls the separator between the line number and text.
// Maps to the -s flag. Default is "\t".
type nlSepFlag string

// NlSep creates a separator flag with the given string.
func NlSep(s string) nlSepFlag { return nlSepFlag(s) }

func (f nlSepFlag) Configure(flags *flags) { flags.sep = string(f) }

// nlStartFlag sets the starting line number (-v flag, default 1).
type nlStartFlag int

// NlStart sets the starting line number.
func NlStart(n int) nlStartFlag { return nlStartFlag(n) }

func (f nlStartFlag) Configure(flags *flags) { v := int(f); flags.start = &v }

// nlIncrementFlag sets the line number increment (-i flag, default 1).
type nlIncrementFlag int

// NlIncrement sets the line number increment.
func NlIncrement(n int) nlIncrementFlag { return nlIncrementFlag(n) }

func (f nlIncrementFlag) Configure(flags *flags) { v := int(f); flags.increment = &v }

// nlWidthFlag sets the field width for line numbers (-w flag, default 6).
type nlWidthFlag int

// NlWidth sets the field width for line numbers.
func NlWidth(n int) nlWidthFlag { return nlWidthFlag(n) }

func (f nlWidthFlag) Configure(flags *flags) { v := int(f); flags.width = &v }

// nlFormatFlag sets the line number format (-n flag).
// Valid values: "ln" (left justified), "rn" (right justified, default), "rz" (right justified zero-padded).
type nlFormatFlag string

const (
	NlFormatLN nlFormatFlag = "ln"
	NlFormatRN nlFormatFlag = "rn"
	NlFormatRZ nlFormatFlag = "rz"
)

// NlFormat sets the line number format.
func NlFormat(f string) nlFormatFlag { return nlFormatFlag(f) }

func (f nlFormatFlag) Configure(flags *flags) { flags.format = f }

type flags struct {
	body      nlBodyFlag
	sep       string
	start     *int
	increment *int
	width     *int
	format    nlFormatFlag
}
