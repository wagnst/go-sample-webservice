package main

type ErrIdInvalid struct {
	message string
}

func NewErrIdInvalid(message string) *ErrIdInvalid {
	return &ErrIdInvalid{
		message: message,
	}
}

func (e *ErrIdInvalid) Error() string {
	return e.message
}

type ErrPlatformInvalid struct {
	message string
}

func NewErrPlatformInvalid(message string) *ErrPlatformInvalid {
	return &ErrPlatformInvalid{
		message: message,
	}
}

func (e *ErrPlatformInvalid) Error() string {
	return e.message
}

type ErrNotFound struct {
	message string
}

func NewErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{
		message: message,
	}
}

func (e *ErrNotFound) Error() string {
	return e.message
}
