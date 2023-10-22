package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/DJohnson2021/go-survey-app/api/controllers"
    "github.com/DJohnson2021/go-survey-app/api/middleware"
)

var RegisteredRoutes = func(app *fiber.App) {

    // Home Route
    app.Get("/", controllers.HomePage)

    // OAuth Routes
    app.Get("/login", controllers.LoginPage)
    app.Get("/api/user/oauth2/google/login", middleware.OauthGoogleLogin)
    app.Get("/api/user/oauth2/google/callback", middleware.OauthGoogleCallBack)

    // User Routes
    // app.Get("/api/user/register", controllers.RegisterUserForm)   // Render registration form
    // app.Post("/api/user/register", controllers.RegisterUser)      // Handle registration submission
    // app.Get("/api/user/signin", controllers.SignInUserForm)       // Render sign-in form
    // app.Post("/api/user/signin", controllers.SignInUser)          // Handle sign-in submission
    //app.Get("/api/user/signout", controllers.SignOutUser)
    app.Get("/api/user/dashboard", middleware.IsUserAuthorized(), controllers.ViewUserDashboard)
    app.Get("/api/user/profile", middleware.IsUserAuthorized(), controllers.ViewUserProfile)
    app.Get("/api/user/survey/results", middleware.IsUserAuthorized(), controllers.ViewUserSurveyResults)

    // Admin Routes
    //app.Get("/api/admin/signin", controllers.SignInAdminForm)     // Render admin sign-in form
    //app.Post("/api/admin/signin", controllers.SignInAdmin)        // Handle admin sign-in submission
    //app.Get("/api/admin/signout", controllers.SignOutAdmin)
    app.Get("/api/admin/dashboard", middleware.IsAdminAuthorized(), controllers.ViewAdminDashboard)
    app.Post("/api/admin/survey/edit", middleware.IsAdminAuthorized(), controllers.EditSurvey)    // Assuming editing means updating/creating.

    // Survey Routes
    app.Get("/api/survey", middleware.IsUserAuthorized(), controllers.ListSurveys)
    app.Get("/api/survey/view", middleware.IsUserAuthorized(), controllers.ViewSurvey)
    app.Post("/api/survey/submit", middleware.IsUserAuthorized(), controllers.SubmitSurvey)
    app.Get("/api/survey/retake", middleware.IsUserAuthorized(), controllers.RetakeSurvey)
}
