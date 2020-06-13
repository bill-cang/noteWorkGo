package custonErr

const (
	ErrBadParams = 1000
)

var (
	ErrBadParamsBody = New(ErrBadParams,"Bad Params.")
)
