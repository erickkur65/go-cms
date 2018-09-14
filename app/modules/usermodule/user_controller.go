package usermodule

import (
	"fmt"
	"html/template"
	"kudotest/app/modules/rolemodule"
	"kudotest/app/modules/utilitymodule"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/josephspurrier/csrfbanana"
)

// UserController to handle user page
type UserController struct {
	Templates          *template.Template
	UserRepository     UserRepository
	GroupRepository    rolemodule.GroupRepository
	UserHelper         UserHelper
	PasswordHashHelper PasswordHashHelper
	SessionHelper      utilitymodule.SessionHelper
}

// GetProfilePage for display the profile page
func (userController *UserController) GetProfilePage(w http.ResponseWriter, r *http.Request) {
	// Context for data in view template
	type Context struct {
		Users             []Pengguna
		ValidationMessage string
	}
	var userList []Pengguna

	// Load profile view
	profilePage := userController.Templates.Lookup("profile.html")

	// Get all users from database
	users := userController.UserRepository.GetUsers()

	// Assign user role
	for _, user := range users {
		user.Grup = userController.GroupRepository.GetGroupByUserID(user.ID)
		userList = append(userList, user)
	}

	context := Context{
		Users:             userList,
		ValidationMessage: userController.SessionHelper.GetValidationMessage(r)}

	profilePage.Execute(w, context)
}

// GetCreateUserPage for display the create user page
func (userController *UserController) GetCreateUserPage(w http.ResponseWriter, r *http.Request) {
	// Context for data in view template
	type Context struct {
		Groups            []rolemodule.Grup
		ValidationMessage string
		CsrfToken         string
	}

	// Load create user view
	createUserTemplate := userController.Templates.Lookup("create_user.html")

	// Get all groups from database
	groups := userController.GroupRepository.GetGroups()

	// Save csrf token to session for comparison when form submission
	sess := userController.SessionHelper.GetSession(r)
	csrfToken := csrfbanana.Token(w, r, sess)
	userController.SessionHelper.SetCSRFToken(r, w, csrfToken)

	context := Context{
		Groups:            groups,
		ValidationMessage: userController.SessionHelper.GetValidationMessage(r),
		CsrfToken:         csrfToken}

	createUserTemplate.Execute(w, context)
}

// GetEditProfilePage for display the edit profile page
func (userController *UserController) GetEditProfilePage(w http.ResponseWriter, r *http.Request) {
	// Context for data in view template
	type Context struct {
		Groups            []rolemodule.Grup
		User              Pengguna
		UserGroupID       int64
		ValidationMessage string
		CsrfToken         string
	}

	vars := mux.Vars(r)

	// Load edit profile view
	editProfileTemplate := userController.Templates.Lookup("edit_profile.html")

	// Load user id from input type hidden in edit profile
	userID, err := strconv.ParseInt(template.HTMLEscapeString(vars["userId"]), 10, 64)
	if err != nil {
		fmt.Println("error when parse user group id")
	}

	// Get all groups from database
	groups := userController.GroupRepository.GetGroups()

	// Get user group id by given user id
	userGroupID := userController.GroupRepository.GetUserGroupIDByUserID(userID)

	// Get user by given user id
	user := userController.UserRepository.GetUserByID(userID)

	// Save csrf token to session for comparison when form submission
	sess := userController.SessionHelper.GetSession(r)
	csrfToken := csrfbanana.Token(w, r, sess)
	userController.SessionHelper.SetCSRFToken(r, w, csrfToken)

	context := Context{
		Groups:            groups,
		User:              user,
		UserGroupID:       userGroupID,
		ValidationMessage: userController.SessionHelper.GetValidationMessage(r),
		CsrfToken:         csrfToken}

	editProfileTemplate.Execute(w, context)
}

// PostCreateUserPage handle the create user page submission
func (userController *UserController) PostCreateUserPage(w http.ResponseWriter, r *http.Request) {
	var user Pengguna
	r.ParseForm()

	// Prevent csrf attack
	csrfToken := userController.SessionHelper.GetCSRFToken(r)
	if csrfToken != r.FormValue("csrfToken") {
		userController.SessionHelper.SetValidationMessage(r, w, "csrf token salah")
		http.Redirect(w, r, "/user/create", http.StatusFound)
		return
	}

	// Prevent xss attack
	hashPassword, err := userController.PasswordHashHelper.HashPassword(
		template.HTMLEscapeString(r.FormValue("password")))
	if err != nil {
		fmt.Println("error when hash user password")
	}

	age, err := strconv.Atoi(template.HTMLEscapeString(r.FormValue("age")))
	if err != nil {
		fmt.Println("error when parse user age")
	}

	userGroupID, err := strconv.ParseInt(template.HTMLEscapeString(r.FormValue("group")), 10, 64)
	if err != nil {
		fmt.Println("error when parse user group id")
	}

	user.Email = template.HTMLEscapeString(r.FormValue("email"))
	user.NamaDepan = template.HTMLEscapeString(r.FormValue("firstname"))
	user.NamaBelakang = template.HTMLEscapeString(r.FormValue("lastname"))
	user.KataSandi = hashPassword
	user.Umur = age

	// Validate user input data
	validationMessage := userController.UserHelper.ValidateCreateNewUser(user)
	if validationMessage == "Sukses" {
		if userController.UserRepository.IsUserEmailExist(user.Email) {
			validationMessage = "Email sudah ada di database"
		} else {
			userID, err := userController.UserRepository.CreateNewUser(user)
			if err != nil {
				fmt.Println("error when create user")
			}

			err = userController.GroupRepository.InsertUserGroup(userID, userGroupID)
			if err != nil {
				fmt.Println("error when insert user group")
			}

			userController.SessionHelper.SetValidationMessage(r, w, "")
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}
	}

	userController.SessionHelper.SetValidationMessage(r, w, validationMessage)
	http.Redirect(w, r, "/user/create", http.StatusFound)
}

// PostEditProfilePage handle the edit profile page submission
func (userController *UserController) PostEditProfilePage(w http.ResponseWriter, r *http.Request) {
	var user Pengguna
	var err error
	r.ParseForm()

	// Prevent csrf attack
	csrfToken := userController.SessionHelper.GetCSRFToken(r)
	if csrfToken != r.FormValue("csrfToken") {
		userController.SessionHelper.SetValidationMessage(r, w, "csrf token salah")
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	// Prevent xss attack
	hashPassword := ""

	if r.FormValue("new-password") != "" {
		hashPassword, err = userController.PasswordHashHelper.HashPassword(
			template.HTMLEscapeString(r.FormValue("new-password")))
		if err != nil {
			fmt.Println("error when hash user password")
		}
	}

	userID, err := strconv.ParseInt(template.HTMLEscapeString(r.FormValue("userId")), 10, 64)
	if err != nil {
		fmt.Println("error when parse user id")
	}

	age, err := strconv.Atoi(template.HTMLEscapeString(r.FormValue("age")))
	if err != nil {
		fmt.Println("error when parse user age")
	}

	userGroupID, err := strconv.ParseInt(template.HTMLEscapeString(r.FormValue("group")), 10, 64)
	if err != nil {
		fmt.Println("error when parse user group id")
	}

	user.Email = template.HTMLEscapeString(r.FormValue("email"))
	user.KataSandi = hashPassword
	user.NamaDepan = template.HTMLEscapeString(r.FormValue("firstname"))
	user.NamaBelakang = template.HTMLEscapeString(r.FormValue("lastname"))
	user.ID = userID
	user.Umur = age

	// Validate user input data
	validationMessage := userController.UserHelper.ValidateEditProfile(user)
	if validationMessage == "Sukses" {
		err = userController.UserRepository.UpdateUser(user)
		if err != nil {
			fmt.Println("error when update user")
		}

		err = userController.GroupRepository.UpdateUserGroup(userID, userGroupID)
		if err != nil {
			fmt.Println("error when update user group")
		}

		userController.SessionHelper.SetValidationMessage(r, w, "")
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	userController.SessionHelper.SetValidationMessage(r, w, validationMessage)
	http.Redirect(w, r, "/user/edit/profile/"+strconv.Itoa(int(userID)), http.StatusFound)
}
