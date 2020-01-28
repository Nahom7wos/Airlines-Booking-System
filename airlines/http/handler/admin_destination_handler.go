package handler

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type DestinationHandler struct {
	tmpl           *template.Template
	destinationSrv flight.DestinationService
}

// NewDestinationHandler initializes and returns new DestinationHandler
func NewDestinationHandler(T *template.Template, dSrv flight.DestinationService) *DestinationHandler {
	return &DestinationHandler{tmpl: T, destinationSrv: dSrv}
}

// Destination displays all the destinations in the database
func (dh *DestinationHandler) Destination(w http.ResponseWriter, r *http.Request) {
	destinations, errs := dh.destinationSrv.Destinations()
	if errs != nil {
		panic(errs)
	}
	dh.tmpl.ExecuteTemplate(w, "admin.destination.layout", destinations)
}

// StoreDestination creates new destination in the database
func (dh *DestinationHandler) StoreDestination(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		dstn := &entity.Destination{}
		price, _ := strconv.Atoi(r.FormValue("destinationPrice"))
		dstn.Name = r.FormValue("destinationName")
		dstn.Price = uint(price)
		dstn.Description = r.FormValue("destinationDescription")
		mf, fh, err := r.FormFile("destinationImage")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		dstn.Image = fh.Filename
		writeFile(&mf, fh.Filename)

		_, errs := dh.destinationSrv.StoreDestination(dstn)

		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/admin/destination", http.StatusSeeOther)

	} else {

		dh.tmpl.ExecuteTemplate(w, "admin.destination.create.layout", nil)

	}
}

// UpdateDestination handle requests on /admin/destination/update
func (dh *DestinationHandler) UpdateDestination(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		dstn, errs := dh.destinationSrv.Destination(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}

		dh.tmpl.ExecuteTemplate(w, "admin.destination.update.layout", dstn)

	} else if r.Method == http.MethodPost {

		dstn := &entity.Destination{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		price, _ := strconv.Atoi(r.FormValue("destinationPrice"))

		dstn.ID = uint(id)
		dstn.Name = r.FormValue("destinationName")
		dstn.Price = uint(price)
		dstn.Image = r.FormValue("image")
		dstn.Description = r.FormValue("destinationDescription")

		mf, fh, err := r.FormFile("destinationImage")
		iferr == nil {
			dstn.Image = fh.Filename
			err = writeFile(&mf, dstn.Image)
		}
		if mf != nil {
			defer mf.Close()
		}
		
		_, errs := dh.destinationSrv.UpdateDestination(dstn)
		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/admin/destination", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/destination", http.StatusSeeOther)
	}

}

// DeleteDestination handles requests on route /admin/destination/delete
func (dh *DestinationHandler) DeleteDestination(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		_, errs := dh.destinationSrv.DeleteDestination(uint(id))
		if len(errs) > 0 {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/destination", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "../../ui", "assets", "images", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
