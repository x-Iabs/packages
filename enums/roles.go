package enums

// Role represents a user role in the system.
type Role string

const (
	// Admin represents the role of an administrator.
	Admin Role = "admin"

	// User represents the role of a regular user.
	User Role = "user"

	// Author represents the role of an author.
	Author Role = "author"

	// Editor represents the role of an editor.
	Editor Role = "editor"
)

// USERS is a list of User and Admin roles for use in the application.
var USERS_GROUP = []Role{Admin, User}

// EDITORS_GROUP is a list of Editor and Admin roles for use in the application.
var EDITORS_GROUP = []Role{Admin, Editor}

// AUTHORS_GROUP is a list of Author and Admin roles for use in the application.
var AUTHORS_GROUP = []Role{Admin, Author}

// ALL_ROLES is a list of all roles for use in the application.
var ALL_ROLES = []Role{Admin, User, Author, Editor}
