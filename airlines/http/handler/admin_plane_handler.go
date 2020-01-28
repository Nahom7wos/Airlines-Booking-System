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
func (ph *PlaneHandler) Plane(w http.ResponseWriter, r *http.Request) {
	planes, errs := ph.planeSrv.Planes()
	if errs != nil {
		panic(errs)
	}
	ph.tmpl.ExecuteTemplate(w, "admin.plane.layout", planes)
}

// StorePlane creates new plane in the database
func (ph *PlaneHandler) StorePlane(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pln := &entity.Plane{}
		capacity, _ := strconv.Atoi(r.FormValue("planeCapacity"))
		status, _ := strconv.ParseBool(r.FormValue("planeStatus"))

		pln.Name = r.FormValue("planeName")
		pln.Capacity = uint(capacity)
		pln.Status = status

		_, errs := ph.planeSrv.StorePlane(pln)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w, r, "/admin/plane", http.StatusSeeOther)

	} else {

		ph.tmpl.ExecuteTemplate(w, "admin.plane.create.layout", nil)

	}
}

// UpdatePlane handle requests on /admin/plane/update
func (ph *PlaneHandler) UpdatePlane(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		pln, errs := ph.planeSrv.Plane(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}

		ph.tmpl.ExecuteTemplate(w, "admin.plane.update.layout", pln)

	} else if r.Method == http.MethodPost {

		pln := &entity.Plane{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		capacity, _ := strconv.Atoi(r.FormValue("planeCapacity"))
		status, _ := strconv.ParseBool(r.FormValue("planeStatus"))

		pln.ID = uint(id)
		pln.Name = r.FormValue("planeName")
		pln.Capacity = uint(capacity)
		pln.Status = status
		_, errs := ph.planeSrv.UpdatePlane(pln)
		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/admin/plane", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/plane", http.StatusSeeOther)
	}

}

// DeletePlane handles requests on route /admin/plane/delete
func (ph *PlaneHandler) DeletePlane(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		_, errs := ph.planeSrv.DeletePlane(uint(id))
		if len(errs) > 0 {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/plane", http.StatusSeeOther)
}

