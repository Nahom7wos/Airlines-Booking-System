package main

import (
	"html/template"
	"net/http"

	"github.com/Nahom7wos/Airlines-Booking-System/airlines/http/handler"
	brepim "github.com/Nahom7wos/Airlines-Booking-System/book/repository"
	bsrvim "github.com/Nahom7wos/Airlines-Booking-System/book/service"
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	frepim "github.com/Nahom7wos/Airlines-Booking-System/flight/repository"
	fsrvim "github.com/Nahom7wos/Airlines-Booking-System/flight/service"
	"github.com/Nahom7wos/Airlines-Booking-System/rtoken"
	urepim "github.com/Nahom7wos/Airlines-Booking-System/user/repository"
	usrvim "github.com/Nahom7wos/Airlines-Booking-System/user/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTables(dbconn *gorm.DB) []error {
	// errs := dbconn.CreateTable(&entity.Destination{}, &entity.Plane{}, &entity.Flight{}).GetErrors()
	// errs := dbconn.CreateTable(&entity.Ticket{}, &entity.User{}).GetErrors()
	errs := dbconn.CreateTable(&entity.Login{}, &entity.Loyalty{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))

	dbconn, err := gorm.Open("postgres", "postgres://postgres:Postgre_1@localhost/airlinesdb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	createTables(dbconn)
	defer dbconn.Close()

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	destinationRepo := frepim.NewDestinationGormRepo(dbconn)
	destinationServ := fsrvim.NewDestinationService(destinationRepo)
	planeRepo := frepim.NewPlaneGormRepo(dbconn)
	planeServ := fsrvim.NewPlaneService(planeRepo)
	flightRepo := frepim.NewFlightGormRepo(dbconn)
	flightServ := fsrvim.NewFlightService(flightRepo)

	ticketRepo := brepim.NewTicketGormRepo(dbconn)
	ticketServ := bsrvim.NewTicketService(ticketRepo)
	userRepo := urepim.NewUserGormRepo(dbconn)
	userServ := usrvim.NewUserService(userRepo)
	sessionRepo := urepim.NewSessionGormRepo(dbconn)
	sessionServ := usrvim.NewSessionService(sessionRepo)
	roleRepo := urepim.NewRoleGormRepo(dbconn)
	roleServ := usrvim.NewRoleService(roleRepo)
	loginRepo := urepim.NewLoginGormRepo(dbconn)
	loginServ := usrvim.NewLoginService(loginRepo)

	sess := configSess()
	uh := handler.NewUserHandler(tmpl, userServ, loginServ, sessionServ, roleServ, sess, csrfSignKey)
	auh := handler.NewAdminUserHandler(tmpl, userServ, loginServ,roleServ, csrfSignKey)
	mainHandler := handler.NewMainHandler(tmpl, destinationServ, flightServ, ticketServ, userServ)
	ticketHandler := handler.NewTicketHandler(tmpl, ticketServ)
	destinationHandler := handler.NewDestinationHandler(tmpl, destinationServ)
	planeHandler := handler.NewPlaneHandler(tmpl, planeServ)
	flightHandler := handler.NewFlightHandler(tmpl, planeServ, destinationServ, flightServ)
	myFlightHandler := handler.NewMyFlightHandler(tmpl, ticketServ, flightServ, userServ)
	roleHandler := handler.NewRoleHandler(tmpl, roleServ)


	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", mainHandler.Index)
	mux.HandleFunc("/book", mainHandler.Book)
	mux.HandleFunc("/checkin", mainHandler.Checkin)
	mux.Handle("/loyalty", uh.Authenticated(http.HandlerFunc(mainHandler.Loyalty)))
	mux.Handle("/myflight", uh.Authenticated(http.HandlerFunc(mainHandler.MyFlight)))
	mux.Handle("/myflight/update", uh.Authenticated(http.HandlerFunc(myFlightHandler.UpdateMyFlight)))
	mux.HandleFunc("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(mainHandler.Admin))))
	http.HandleFunc("/login", uh.Login)
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/signup", uh.Signup)

	//admin paths
	mux.Handle("/admin/flight", uh.Authenticated(uh.Authorized(http.HandlerFunc(flightHandler.Flight))))
	mux.Handle("/admin/flight/create", uh.Authenticated(uh.Authorized(http.HandlerFunc(flightHandler.StoreFlight))))
	mux.Handle("/admin/flight/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(flightHandler.UpdateFlight))))
	mux.Handle("/admin/flight/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(flightHandler.DeleteFlight))))
	mux.Handle("/admin/destination", uh.Authenticated(uh.Authorized(http.HandlerFunc(destinationHandler.Destination))))
	mux.Handle("/admin/destination/create", uh.Authenticated(uh.Authorized(http.HandlerFunc(destinationHandler.StoreDestination))))
	mux.Handle("/admin/destination/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(destinationHandler.UpdateDestination))))
	mux.Handle("/admin/destination/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(destinationHandler.DeleteDestination))))
	mux.Handle("/admin/plane", uh.Authenticated(uh.Authorized(http.HandlerFunc(planeHandler.Plane))))
	mux.Handle("/admin/plane/create", uh.Authenticated(uh.Authorized(http.HandlerFunc(planeHandler.StorePlane))))
	mux.Handle("/admin/plane/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(planeHandler.UpdatePlane))))
	mux.Handle("/admin/plane/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(planeHandler.DeletePlane))))
	mux.Handle("/admin/ticket", uh.Authenticated(uh.Authorized(http.HandlerFunc(ticketHandler.Ticket))))
	mux.Handle("/admin/ticket/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(ticketHandler.UpdateTicket))))
	mux.Handle("/admin/ticket/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(ticketHandler.DeleteTicket))))
	mux.Handle("/admin/user", uh.Authenticated(uh.Authorized(http.HandlerFunc(auh.AdminUser))))
	mux.Handle("/admin/user/create", uh.Authenticated(uh.Authorized(http.HandlerFunc(auh.AdminStoreUser))))
	mux.Handle("/admin/user/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(auh.AdminUpdateUser))))
	mux.Handle("/admin/user/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(auh.AdminDeleteUser))))
	mux.Handle("/admin/role", uh.Authenticated(uh.Authorized(http.HandlerFunc(roleHandler.role))))
	mux.Handle("/admin/role/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(roleHandler.UpdateRole))))
	mux.Handle("/admin/role/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(roleHandler.DeleteRole))))

	
	http.ListenAndServe(":8080", mux)

}
