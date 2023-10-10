package routes

import (
	"github.com/DJohnson2021/go-survey-app/api/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {

	// User Routes
	r.HandleFunc("/api/user/register", controllers.RegisterUserForm).Methods("GET") // Render registration form
	r.HandleFunc("/api/user/register", controllers.RegisterUser).Methods("POST")    // Handle registration submission
	r.HandleFunc("/api/user/signin", controllers.SignInUserForm).Methods("GET")     // Render sign-in form
	r.HandleFunc("/api/user/signin", controllers.SignInUser).Methods("POST")        // Handle sign-in submission
	r.HandleFunc("/api/user/signout", controllers.SignOutUser).Methods("GET")
	r.HandleFunc("/api/user/profile", controllers.ViewUserProfile).Methods("GET")
	r.HandleFunc("/api/user/survey/results", controllers.ViewUserSurveyResults).Methods("GET")

	// Admin Routes
	r.HandleFunc("/api/admin/signin", controllers.SignInAdminForm).Methods("GET") // Render admin sign-in form
	r.HandleFunc("/api/admin/signin", controllers.SignInAdmin).Methods("POST")    // Handle admin sign-in submission
	r.HandleFunc("/api/admin/signout", controllers.SignOutAdmin).Methods("GET")
	r.HandleFunc("/api/admin/dashboard", controllers.ViewAdminDashboard).Methods("GET")
	r.HandleFunc("/api/admin/survey/edit", controllers.EditSurvey).Methods("POST") // Assuming editing means updating/creating.

	// Survey Routes
	r.HandleFunc("/api/survey", controllers.ListSurveys).Methods("GET")
	r.HandleFunc("/api/survey/view", controllers.ViewSurvey).Methods("GET")
	r.HandleFunc("/api/survey/submit", controllers.SubmitSurvey).Methods("POST")
}
