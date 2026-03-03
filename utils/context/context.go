package context

type SessionUser struct {
	UserID    string
	Email     string
	FirstName string
	LastName  string
}

var sessionUser SessionUser = SessionUser{}

func SetSessionUser(user SessionUser) {
	sessionUser = user
}

func GetSessionUser() *SessionUser {
	return &sessionUser
}
