package app

type Role interface {
	IsMetBy(role Role) bool
}

type AdminRole struct {
}

func (r *AdminRole) IsMetBy(role Role) bool {
	switch role.(type) {
	case *AdminRole:
		return true
	default:
		return false
	}
}

type UserRole struct {
}

func (r *UserRole) IsMetBy(role Role) bool {
	return true
}
