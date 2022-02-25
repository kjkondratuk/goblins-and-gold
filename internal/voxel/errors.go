package voxel

type ErrorAlreadyExists struct {
	val string
}

func NewErrorAlreadyExists(val string) error {
	return &ErrorAlreadyExists{
		val: val,
	}
}

func (e *ErrorAlreadyExists) Error() string {
	return e.val + " already exists!"
}
