package utilitymodule

import (
	"fmt"
	"kudotest/app/modules/permissionmodule"
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionHelper helper to handle session
type SessionHelper struct {
	Store                *sessions.CookieStore
	PermissionRepository permissionmodule.PermissionRepository
}

// GetSession get session object
func (sessionHelper *SessionHelper) GetSession(r *http.Request) *sessions.Session {
	session, _ := sessionHelper.Store.Get(r, "session")
	return session
}

// IsLoggedIn check whether the user login or not
func (sessionHelper *SessionHelper) IsLoggedIn(r *http.Request) bool {
	sess, err := sessionHelper.Store.Get(r, "session")

	if err == nil && sess.Values["is_loggedin"] == "true" {
		return true
	}

	return false
}

// GetUserGroupPermissions get user group permission from email in session
func (sessionHelper *SessionHelper) GetUserGroupPermissions(r *http.Request) []string {
	sess, err := sessionHelper.Store.Get(r, "session")

	if err != nil {
		fmt.Println("error when get session")
	}

	email := sess.Values["email"].(string)
	userGroupPermissions := sessionHelper.PermissionRepository.GetGroupPermissions(email)

	return userGroupPermissions
}

// SetValidationMessage for set validation message
func (sessionHelper *SessionHelper) SetValidationMessage(r *http.Request, w http.ResponseWriter, message string) {
	sess, err := sessionHelper.Store.Get(r, "session")

	if err != nil {
		fmt.Println("error when get session(set validation message)")
	}

	if message != "Sukses" {
		sess.Values["validation_message"] = message
		sess.Save(r, w)
	}
}

// GetValidationMessage for get validation message
func (sessionHelper *SessionHelper) GetValidationMessage(r *http.Request) string {
	sess, err := sessionHelper.Store.Get(r, "session")
	validationMessage := ""

	if err != nil {
		fmt.Println("error when get session(get validation message)")
	}

	if sess.Values["validation_message"] != nil && sess.Values["validation_message"] != "" {
		validationMessage = sess.Values["validation_message"].(string)
	}

	return validationMessage
}

// SetCSRFToken set csrf token
func (sessionHelper *SessionHelper) SetCSRFToken(r *http.Request, w http.ResponseWriter, token string) {
	sess, err := sessionHelper.Store.Get(r, "session")

	if err != nil {
		fmt.Println("error when get session(set csrf token)")
	}

	sess.Values["csrf_token"] = token
	sess.Save(r, w)
}

// GetCSRFToken get csrf token
func (sessionHelper *SessionHelper) GetCSRFToken(r *http.Request) string {
	sess, err := sessionHelper.Store.Get(r, "session")
	token := ""

	if err != nil {
		fmt.Println("error when get session(get csrf token)")
	}

	if sess.Values["csrf_token"] != nil && sess.Values["csrf_token"] != "" {
		token = sess.Values["csrf_token"].(string)
	}

	return token
}
