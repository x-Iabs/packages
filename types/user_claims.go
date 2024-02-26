package types

import (
	"github.com/golang-jwt/jwt"
	"github.com/x-Iabs/packages/enums"
)

// UserClaims represents the claims of a user in the system.
type UserClaims struct {
	ID       string       `json:"id"`
	Username string       `json:"username"`
	Roles    []enums.Role `json:"roles"`
	jwt.StandardClaims
}

// IsAdmin checks if the user has the admin role.
func (uc *UserClaims) IsAdmin() bool {
	for _, role := range uc.Roles {
		if role == enums.Admin {
			return true
		}
	}
	return false
}

// IsUser checks if the user has the user role.
func (uc *UserClaims) IsUser() bool {
	for _, role := range uc.Roles {
		if role == enums.User {
			return true
		}
	}
	return false
}

// IsAuthor checks if the user has the author role.
func (uc *UserClaims) IsAuthor() bool {
	for _, role := range uc.Roles {
		if role == enums.Author {
			return true
		}
	}
	return false
}

// IsEditor checks if the user has the editor role.
func (uc *UserClaims) IsEditor() bool {
	for _, role := range uc.Roles {
		if role == enums.Editor {
			return true
		}
	}
	return false
}
