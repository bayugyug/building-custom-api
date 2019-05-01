package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
	"github.com/bayugyug/building-custom-api/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// APIEndpoints the end-points-url mapping
type APIEndpoints interface {
	BuildCreate(w http.ResponseWriter, r *http.Request)
	BuildingUpdate(w http.ResponseWriter, r *http.Request)
	BuildingGet(w http.ResponseWriter, r *http.Request)
	BuildingGetOne(w http.ResponseWriter, r *http.Request)
	BuildingDelete(w http.ResponseWriter, r *http.Request)
}

// Response is the reply object
type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result,omitempty"`
}

// Building the api handler
type Building struct {
	Context context.Context
	Storage *drivers.Storage
}

// Welcome index page
func (b *Building) Welcome(w http.ResponseWriter, r *http.Request) {
	//good
	render.JSON(w, r, Response{
		Status: "Welcome!",
	})
}

// BuildCreate save a row in store
func (b *Building) BuildCreate(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingCreate()
	//sanity check
	if err := render.Bind(r, data); err != nil {
		//206
		b.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}
	pid, err := data.Create(b.Context, b.Storage)
	//chk
	if err != nil {
		switch err {
		case models.ErrRecordExists:
			//409
			b.ReplyErrContent(w, r, http.StatusConflict, err.Error())
		case models.ErrMissingRequiredParameters:
			//206
			b.ReplyErrContent(w, r, http.StatusPartialContent, err.Error())
		default:
			//500
			b.ReplyErrContent(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	//good
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, Response{
		Status: "Success",
		Result: pid,
	})
}

// BuildingUpdate update row in store
func (b *Building) BuildingUpdate(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingUpdate()
	//sanity check
	if err := render.Bind(r, data); err != nil {
		//206
		b.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}
	//check
	if err := data.Update(b.Context, b.Storage); err != nil {
		switch err {
		case models.ErrRecordMismatch:
			//409
			b.ReplyErrContent(w, r, http.StatusConflict, err.Error())
		case models.ErrMissingRequiredParameters:
			//206
			b.ReplyErrContent(w, r, http.StatusPartialContent, err.Error())
		case models.ErrRecordNotFound:
			//204 or 404?
			b.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		default:
			//500
			b.ReplyErrContent(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	//good
	render.JSON(w, r, Response{
		Status: "Success",
	})
}

// BuildingGet list all
func (b *Building) BuildingGet(w http.ResponseWriter, r *http.Request) {
	data := &models.BuildingGetParams{}
	//check
	rows, err := data.GetAll(b.Context, b.Storage)
	//chk
	if err != nil {
		//404
		b.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		return
	}
	//good
	render.JSON(w, r, Response{
		Status: "Success",
		Result: rows,
	})
}

// BuildingGetOne get 1 row per id
func (b *Building) BuildingGetOne(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingGetOne(strings.TrimSpace(chi.URLParam(r, "id")))
	//chk
	if data.ID == "" {
		//206
		b.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}
	//check
	row, err := data.Get(b.Context, b.Storage)
	//chk
	if err != nil {
		//404
		b.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		return
	}
	//good
	render.JSON(w, r, Response{
		Status: "Success",
		Result: row,
	})
}

// BuildingDelete remove from store
func (b *Building) BuildingDelete(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingDelete(strings.TrimSpace(chi.URLParam(r, "id")))
	//chk
	if data.ID == "" {
		//206
		b.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}
	//chk
	if err := data.Remove(b.Context, b.Storage); err != nil {
		switch err {
		case models.ErrRecordNotFound:
			//404
			b.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		default:
			//500
			b.ReplyErrContent(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	//good
	render.JSON(w, r, Response{
		Status: "Success",
	})
}

// ReplyErrContent send err-code/err-msg
func (b *Building) ReplyErrContent(w http.ResponseWriter, r *http.Request, code int, msg string) {
	render.Status(r, code)
	render.JSON(w, r, Response{
		Status: msg,
	})
}

// HealthCheck index page
func (b *Building) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//good
	render.JSON(w, r, Response{
		Status: "Building API Service: " + time.Now().Format(time.RFC3339),
	})
}
