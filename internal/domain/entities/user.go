package entities

import "github.com/google/uuid"

type Role struct {
	Name  string
	Level int
}

var (
	RoleUser  Role            = Role{"user", 0}
	RoleAdmin Role            = Role{"admin", 1}
	Roles     map[string]Role = rolesToMap(RoleUser, RoleAdmin)
)

func rolesToMap(roles ...Role) map[string]Role {
	m := make(map[string]Role)
	for _, r := range roles {
		m[r.Name] = r
	}
	return m
}

// retuns role struct from name, defaults to RoleUser
func RoleFromString(roleName string) Role {
	r, ok := Roles[roleName]
	if !ok {
		return RoleUser
	}
	return r
}

// ex. IsRoleHigher("user", "admin") -> false, user = 0, admin = 1
func IsRoleHigherOrEqual(role string, thanRole string) bool {
	r1 := RoleFromString(role)
	r2 := RoleFromString(role)

	return r1.Level >= r2.Level
}

type Email string

type User struct {
	Name         string
	Email        Email
	Role         Role
	PasswordHash string
	List         []uuid.UUID
}

func ConstructUser(email Email, name, passwordHash string, role Role, list []uuid.UUID) *User {
	return &User{
		Email:        email,
		Name:         name,
		Role:         role,
		PasswordHash: passwordHash,
		List:         []uuid.UUID{},
	}
}

type UserOption func(*User)

func WithRole(role Role) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func NewUser(email Email, name, passwordHash string, opts ...UserOption) (*User, error) {
	// TODO: validation

	user := ConstructUser(email, name, passwordHash, RoleUser, []uuid.UUID{})

	for _, o := range opts {
		o(user)
	}

	return user, nil
}
