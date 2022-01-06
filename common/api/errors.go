package api

const (
	RegisterUserAlreadyExists uint32 = iota + 1
	RegisterUserNameEmpty
	RegisterPasswordEmpty
	LoginUserNotFound
	LoginWrongPassword

	InternalError uint32 = 1000
)
