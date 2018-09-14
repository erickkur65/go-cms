package main

import (
	"fmt"
	"kudotest/app/modules/authmodule"
	"kudotest/app/modules/permissionmodule"
	"kudotest/app/modules/rolemodule"
	"kudotest/app/modules/usermodule"
	"kudotest/app/modules/utilitymodule"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	// Define database configuration
	databaseConfig := utilitymodule.DatabaseConfig{
		Username: "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     "3306",
		DbName:   "kudotest_db"}
	databaseHelper := utilitymodule.DatabaseHelper{DatabaseConfig: databaseConfig}
	dbConnection := databaseHelper.Connect()

	// Define template for view
	templateHelper := utilitymodule.TemplateHelper{}
	templates := templateHelper.PopulateTemplates()

	// Define model repository
	groupRepository := rolemodule.GroupRepository{Database: dbConnection}
	permissionRepository := permissionmodule.PermissionRepository{Database: dbConnection}
	userRepository := usermodule.UserRepository{Database: dbConnection}

	// Define helper
	passwordHashHelper := usermodule.PasswordHashHelper{}
	userHelper := usermodule.UserHelper{}
	authHelper := authmodule.AuthHelper{
		UserRepository:  userRepository,
		GroupRepository: groupRepository}
	sessionHelper := utilitymodule.SessionHelper{
		Store:                sessions.NewCookieStore([]byte("secret-pass")),
		PermissionRepository: permissionRepository}
	middlewareHelper := utilitymodule.MiddlewareHelper{
		SessionHelper: sessionHelper}

	// Define controller
	userController := usermodule.UserController{
		Templates:          templates,
		UserRepository:     userRepository,
		GroupRepository:    groupRepository,
		UserHelper:         userHelper,
		PasswordHashHelper: passwordHashHelper,
		SessionHelper:      sessionHelper}
	authController := authmodule.AuthController{
		Templates:     templates,
		AuthHelper:    authHelper,
		SessionHelper: sessionHelper}

	// Define router for web
	r := mux.NewRouter()
	r.HandleFunc("/user/create", middlewareHelper.CheckLoginAndPermission(userController.GetCreateUserPage)).Methods("GET")
	r.HandleFunc("/user/create", middlewareHelper.CheckLoginAndPermission(userController.PostCreateUserPage)).Methods("POST")
	r.HandleFunc("/user/edit/profile/{userId}", middlewareHelper.CheckLoginAndPermission(userController.GetEditProfilePage)).Methods("GET")
	r.HandleFunc("/user/edit/profile", middlewareHelper.CheckLoginAndPermission(userController.PostEditProfilePage)).Methods("POST")
	r.HandleFunc("/profile", middlewareHelper.CheckLoginAndPermission(userController.GetProfilePage)).Methods("GET")
	r.HandleFunc("/logout", middlewareHelper.CheckLoginAndPermission(authController.GetLogout)).Methods("GET")
	r.HandleFunc("/", authController.GetLoginPage).Methods("GET")
	r.HandleFunc("/login", authController.PostLoginPage).Methods("POST")
	http.Handle("/", r)
	http.Handle("/static/", http.FileServer(http.Dir("public")))

	fmt.Println("running server on 8081")
	err := http.ListenAndServe("0.0.0.0:8081", nil)

	if err != nil {
		fmt.Println("error in main.go: ", err)
	}
}
