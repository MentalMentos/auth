package extended

import (
	"context"
	"example.com/m/internal/model"
)

type Role int

const (
	RoleAdmin Role = iota
	RoleUser
)

// extended_methods.go

func IsAdmin(ctx context.Context, user *model.User) (Role, error) {
	// Your user authentication logic here
	// Make sure the password check is correct before proceeding
	// After successful authentication, determine the user's role.
	if user.Role == 1 {
		return RoleAdmin, nil
	}
	return RoleUser, nil
}
