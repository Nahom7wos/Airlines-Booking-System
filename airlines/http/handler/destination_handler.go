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

// DestinationStore creates new destination in the database
func (dh *DestinationHandler) DestinationStore(w http.ResponseWriter, r *http.Request) {
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
func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
