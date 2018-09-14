package authmodule

import (
	"html/template"
	"kudotest/app/modules/usermodule"
	"kudotest/app/modules/utilitymodule"
	"net/http"

	"github.com/josephspurrier/csrfbanana"
)

// AuthController to handle authentication page
type AuthController struct {
	Templates     *template.Template
	AuthHelper    AuthHelper
	SessionHelper utilitymodule.SessionHelper
}

// GetLoginPage to render login page
func (authController *AuthController) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	// Context for data in view template
	type Context struct {
		ValidationMessage string
		CsrfToken         string
	}

	// Load login view
	loginTemplate := authController.Templates.Lookup("login.html")

	// Save csrf token to session for comparison when form submission
	sess := authController.SessionHelper.GetSession(r)
	csrfToken := csrfbanana.Token(w, r, sess)
	authController.SessionHelper.SetCSRFToken(r, w, csrfToken)

	context := Context{
		ValidationMessage: authController.SessionHelper.GetValidationMessage(r),
		CsrfToken:         csrfToken}

	loginTemplate.Execute(w, context)
}

// PostLoginPage to handle form submission in login page
func (authController *AuthController) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	var user usermodule.Pengguna
	r.ParseForm()

	// Prevent csrf attack
	csrfToken := authController.SessionHelper.GetCSRFToken(r)
	if csrfToken != r.FormValue("csrfToken") {
		authController.SessionHelper.SetValidationMessage(r, w, "csrf token salah")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Prevent xss attack
	user.Email = template.HTMLEscapeString(r.FormValue("email"))
	user.KataSandi = template.HTMLEscapeString(r.FormValue("password"))
	validationMessage := authController.AuthHelper.ValidateLoginData(user)

	if validationMessage == "Sukses" {
		// Set session for user
		sess := authController.SessionHelper.GetSession(r)
		sess.Values["is_loggedin"] = "true"
		sess.Values["email"] = user.Email
		sess.Values["validation_message"] = ""
		sess.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	// Set validation message for validation fail
	authController.SessionHelper.SetValidationMessage(r, w, validationMessage)
	http.Redirect(w, r, "/", http.StatusFound)
}

// GetLogout handle user logout
func (authController *AuthController) GetLogout(w http.ResponseWriter, r *http.Request) {
	sess := authController.SessionHelper.GetSession(r)

	// Set session logged in to false
	if sess.Values["is_loggedin"] == "true" {
		sess.Values["is_loggedin"] = "false"
		sess.Save(r, w)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
