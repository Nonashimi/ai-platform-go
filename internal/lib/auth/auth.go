package auth

import "net/http"

const (
	CtxUserID = "user_id"
	CtxRole   = "role"
)

func GetUserID(r *http.Request) (uint, bool) {
	val := r.Context().Value(CtxUserID)
	if val == nil {
		return 0, false
	}

	id, ok := val.(uint)
	if !ok {
		return 0, false
	}

	return id, true
}

func GetRole(r *http.Request) (string, bool) {
	val := r.Context().Value(CtxRole)
	if val == nil {
		return "", false
	}
	role, ok := val.(string)
	return role, ok
}
