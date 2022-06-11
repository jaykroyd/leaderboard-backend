package app

type User interface {
	Identifier() string
}

type HubUser struct {
	user  string
	email string
}

func (u *HubUser) Identifier() string {
	return u.email
}
