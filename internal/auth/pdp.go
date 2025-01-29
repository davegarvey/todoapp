package auth

import (
	"todo_app/internal/store"
)

func EvaluatePolicy(req AuthZENRequest) bool {
	user, exists := store.UserStore[req.Subject.ID]
	if !exists {
		return false
	}

	switch req.Action.Name {
	case "can_read_user":
		return true
	case "can_read_todos":
		return true
	case "can_create_todo":
		return hasAnyRole(user, []string{"admin", "editor"})
	case "can_update_todo":
		if hasRole(user, "evil_genius") {
			return true
		}
		return hasRole(user, "editor") && user.Email == req.Resource.Properties.OwnerID
	case "can_delete_todo":
		if hasRole(user, "admin") {
			return true
		}
		return hasRole(user, "editor") && user.Email == req.Resource.Properties.OwnerID
	default:
		return false
	}
}

func hasRole(user store.User, role string) bool {
	for _, r := range user.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func hasAnyRole(user store.User, roles []string) bool {
	for _, role := range roles {
		if hasRole(user, role) {
			return true
		}
	}
	return false
}
