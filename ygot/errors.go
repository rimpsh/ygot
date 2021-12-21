package ygot

// UnsupportedError allows a user to ignore currently unsupported types which
// are defined within used YANG files.
type UnsupportedError struct {
	errMessage string
}

func (e UnsupportedError) Error() string {
	return e.errMessage
}
