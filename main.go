/// Routes

package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/s0lica/BitsLab/dbbuilder"
	"github.com/s0lica/BitsLab/internal/api"
	"github.com/s0lica/BitsLab/internal/auth"
	"github.com/s0lica/BitsLab/routes"
)

func main() {
	//STYLESHEETS + SCRIPTS
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("./stylesheets"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	//ROUTING NO LOGIN
	http.HandleFunc("/login", routes.Login)
	http.HandleFunc("/sign_up", routes.Sign_up)
	http.HandleFunc("/Login_user", auth.Login_user)
	http.HandleFunc("/Sign_up_user", auth.Sign_up_user)
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/probleme", routes.Problem)
	http.HandleFunc("/logouthandle", auth.LogoutHandler)
	//DYNAMIC ROUTING
	http.HandleFunc("/problems/{id}", routes.ProblemHandler)
	//ROUTING AUTHREQUIRED
	http.HandleFunc("/user", auth.AuthRequired(routes.Userpage))
	//ROUTING ADMINREQUIRED
	http.HandleFunc("/admin/create_problem", auth.AdminRequired(routes.Create_problem))
	//ACTION ROUTING
	http.HandleFunc("/Create_problem", api.Create_problemHandler)
	//TABLES BUILDER
	dbbuilder.Build_databases()
	//sv start
	http.ListenAndServe(":8000", nil)
}
