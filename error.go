//go:generate stringer -type=Error

package spec2test

type Error int

const (
	ErrorStructTypeRequired Error = iota
	ErrorResourceAlreadyAdded
	ErrorContentTypeNotSupported
)

func (err Error) Error() string {
	return err.String()
}
