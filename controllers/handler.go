package controllers

import (
	"context"
	"net/http"
	"strings"

	"github.com/bayugyug/building-custom-api/driver"
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

// APIResponse is the reply object
type APIResponse struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// APIHandler the api handler
type APIHandler struct {
	Context context.Context
	Storage *driver.Storage
}

// Welcome index page
func (resp *APIHandler) Welcome(w http.ResponseWriter, r *http.Request) {

	//good
	render.JSON(w, r, APIResponse{
		Code:   200,
		Status: "Welcome!",
	})
}

// BuildCreate save a row in store
func (resp *APIHandler) BuildCreate(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingCreate()

	//sanity check
	if err := render.Bind(r, data); err != nil {
		//206
		resp.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}
	pid, err := data.Create(resp.Context, resp.Storage)
	//chk
	if err != nil {
		//400
		resp.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//good
	render.JSON(w, r, APIResponse{
		Code:   200,
		Status: "Success",
		Result: pid,
	})
}

// BuildingUpdate update row in store
func (resp *APIHandler) BuildingUpdate(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingUpdate()

	//sanity check
	if err := render.Bind(r, data); err != nil {
		//206
		resp.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}

	//chk
	if err := data.Update(resp.Context, resp.Storage); err != nil {
		//400
		resp.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//good
	render.JSON(w, r, APIResponse{
		Code:   200,
		Status: "Success",
	})
}

// BuildingGet list all
func (resp *APIHandler) BuildingGet(w http.ResponseWriter, r *http.Request) {
	data := &models.BuildingGetOneParams{}

	//check
	rows, err := data.GetAll(resp.Context, resp.Storage)

	//chk
	if err != nil {
		//404
		resp.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		return
	}
	//good
	render.JSON(w, r, APIResponse{
		Code:   200,
		Status: "Success",
		Result: rows,
	})
}

// BuildingGetOne get 1 row per id
func (resp *APIHandler) BuildingGetOne(w http.ResponseWriter, r *http.Request) {

	data := models.NewBuildingGetOne(strings.TrimSpace(chi.URLParam(r, "id")))

	//chk
	if len(data.ID) == 0 {
		//206
		resp.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}

	//check
	row, err := data.Get(resp.Context, resp.Storage)

	//chk
	if err != nil {
		//404
		resp.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		return
	}

	//good
	render.JSON(w, r, APIResponse{
		Code:   200,
		Status: "Success",
		Result: row,
	})
}

// BuildingDelete remove from store
func (resp *APIHandler) BuildingDelete(w http.ResponseWriter, r *http.Request) {

	data := models.NewBuildingDelete(strings.TrimSpace(chi.URLParam(r, "id")))

	//chk
	if data.ID == "" {
		//206
		resp.ReplyErrContent(w, r, http.StatusPartialContent, http.StatusText(http.StatusPartialContent))
		return
	}

	//chk
	if err := data.Remove(resp.Context, resp.Storage); err != nil {
		//400
		resp.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//good
	render.JSON(w, r, APIResponse{
		Code:   200,
		Status: "Success",
	})
}

// ReplyErrContent send err-code/err-msg
func (resp *APIHandler) ReplyErrContent(w http.ResponseWriter, r *http.Request, code int, msg string) {
	render.Status(r, code)
	render.JSON(w, r, APIResponse{
		Code:   code,
		Status: msg,
	})
}
