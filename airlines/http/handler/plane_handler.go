package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

type PlaneHandler struct {
	tmpl     *template.Template
	planeSrv flight.PlaneService
}

// NewPlaneHandler initializes and returns new PlaneHandler
func NewPlaneHandler(T *template.Template, pSrv flight.PlaneService) *PlaneHandler {
	return &PlaneHandler{tmpl: T, planeSrv: pSrv}
}

// Plane displays all the planes in the database
func (dh *PlaneHandler) Plane(w http.ResponseWriter, r *http.Request) {
	planes, errs := dh.planeSrv.Planes()
	if errs != nil {
		panic(errs)
	}
	dh.tmpl.ExecuteTemplate(w, "admin.plane.layout", planes)
}

// PlaneStore creates new plane in the database
func (dh *PlaneHandler) PlaneStore(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pln := &entity.Plane{}
		capacity, _ := strconv.Atoi(r.FormValue("planeCapacity"))
		pln.Name = r.FormValue("planeName")
		pln.Capacity = uint(capacity)

		http.Redirect(w, r, "/admin/plane", http.StatusSeeOther)

	} else {

		dh.tmpl.ExecuteTemplate(w, "admin.plane.create.layout", nil)

	}
}
