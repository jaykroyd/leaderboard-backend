package app

type RoleIdentifier interface {
	Identify(user User) Role
}

type AdminIdentifier struct {
	roles map[string]int
}

func NewAdminIdentifier(roles map[string]int) *AdminIdentifier {
	return &AdminIdentifier{
		roles: roles,
	}
}

func (i *AdminIdentifier) Identify(user User) Role {
	level, ok := i.roles[user.Identifier()]
	if !ok {
		return new(UserRole)
	}

	switch level {
	default:
		return new(AdminRole)
	}
}
