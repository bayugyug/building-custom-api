package handler

import (
	"context"
	"net/http"
	"strings"

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
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
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
		Code:   200,
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
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//good
	render.JSON(w, r, Response{
		Code:   200,
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

	//chk
	if err := data.Update(b.Context, b.Storage); err != nil {
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//good
	render.JSON(w, r, Response{
		Code:   200,
		Status: "Success",
	})
}

// BuildingGet list all
func (b *Building) BuildingGet(w http.ResponseWriter, r *http.Request) {
	data := &models.BuildingGetOneParams{}

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
		Code:   200,
		Status: "Success",
		Result: rows,
	})
}

// BuildingGetOne get 1 row per id
func (b *Building) BuildingGetOne(w http.ResponseWriter, r *http.Request) {

	data := models.NewBuildingGetOne(strings.TrimSpace(chi.URLParam(r, "id")))

	//chk
	if len(data.ID) == 0 {
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
		Code:   200,
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
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//good
	render.JSON(w, r, Response{
		Code:   200,
		Status: "Success",
	})
}

// ReplyErrContent send err-code/err-msg
func (b *Building) ReplyErrContent(w http.ResponseWriter, r *http.Request, code int, msg string) {
	render.Status(r, code)
	render.JSON(w, r, Response{
		Code:   code,
		Status: msg,
	})
}
