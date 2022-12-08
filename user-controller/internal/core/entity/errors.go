package entity

type Error string

func (err Error) Error() string {
	return string(err)
}

const (
	DatabaseError         = Error("database error")
	NotFoundError         = Error("not found")
	PasswordMismatchError = Error("password mismatch")
	InvalidAuthDataError  = Error("invalid auth data")
	TokenError            = Error("token error")
	UserAlreadyExistError = Error("user already exist error")
	UsersExistsError      = Error("users exists")
)
