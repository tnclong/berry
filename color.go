package berry

// convenient basic foreground colors, useful in many cases.
var (
	Black   = New(FgBlack)
	Red     = New(FgRed)
	Green   = New(FgGreen)
	Yellow  = New(FgYellow)
	Blue    = New(FgBlue)
	Magenta = New(FgMagenta)
	Cyan    = New(FgCyan)
	White   = New(FgWhite)

	// It's convenient global variable when you need controll reset manually
	// for higher performance.
	RReset = New(Reset)
)
