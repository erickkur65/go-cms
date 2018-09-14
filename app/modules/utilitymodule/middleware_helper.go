package utilitymodule

import (
	"net/http"
	"strings"
)

// MiddlewareHelper handle about middleware
type MiddlewareHelper struct {
	SessionHelper SessionHelper
}

const (
	ReadPermission   = "read"
	CreatePermission = "create"
	UpdatePermission = "edit"
)

// CheckLoginAndPermission user must login first to access specific page
func (middlewareHelper *MiddlewareHelper) CheckLoginAndPermission(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !middlewareHelper.SessionHelper.IsLoggedIn(r) {
			errorMessage := "Harus login terlebih dahulu untuk mengakses routing ini"
			middlewareHelper.SessionHelper.SetValidationMessage(r, w, errorMessage)

			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		// Manage permission for user access web routing
		// Super admin role can access profile, creater new user, and update user
		// Admin role can access profile and create new user
		userGroupPermission := middlewareHelper.SessionHelper.GetUserGroupPermissions(r)
		urlPath := r.URL.Path
		urlPermission := ""
		isAllow := false

		if urlPath == "/logout" {
			handler(w, r)
		}

		if urlPath == "/profile" {
			urlPermission = ReadPermission
		} else if urlPath == "/user/create" {
			urlPermission = CreatePermission
		} else if strings.Contains(urlPath, "edit") {
			urlPermission = UpdatePermission
		}

		for _, permission := range userGroupPermission {
			if permission == urlPermission {
				isAllow = true
			}
		}

		if !isAllow {
			errorMessage := "Grup pengguna anda tidak bisa mengakses halaman ini"
			middlewareHelper.SessionHelper.SetValidationMessage(r, w, errorMessage)

			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}

		handler(w, r)
	}
}
