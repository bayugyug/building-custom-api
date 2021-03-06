package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/bayugyug/building-custom-api/configs"
	"github.com/bayugyug/building-custom-api/drivers"
	"github.com/bayugyug/building-custom-api/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// BuildingEndpoints the end-points-url mapping
type BuildingEndpoints interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// Response is the reply object
type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result,omitempty"`
	Total  int         `json:"total,omitempty"`
}

// Building the api handler
type Building struct {
	Storage *drivers.Storage
}

// NewBuilding new instance
func NewBuilding() *Building {
	return &Building{
		Storage: drivers.NewStorage(),
	}
}

// Welcome index page
func (b *Building) Welcome(w http.ResponseWriter, r *http.Request) {
	//good
	render.JSON(w, r, Response{
		Status: "Welcome!",
	})
}

// Create save a row in store
func (b *Building) Create(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingCreate()
	//sanity check
	if err := render.Bind(r, data); err != nil {
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	pid, err := data.Create(b.Storage)
	//chk
	if err != nil {
		switch err {
		case models.ErrRecordExists:
			//409
			b.ReplyErrContent(w, r, http.StatusConflict, err.Error())
		case models.ErrMissingRequiredParameters:
			//400
			b.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
		default:
			//500
			b.ReplyErrContent(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	//good
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, Response{
		Status: "success",
		Result: pid,
	})
}

// Update update row in store
func (b *Building) Update(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingUpdate()
	//sanity check
	if err := render.Bind(r, data); err != nil {
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	//check
	if err := data.Update(b.Storage); err != nil {
		switch err {
		case models.ErrRecordMismatch:
			//409
			b.ReplyErrContent(w, r, http.StatusConflict, err.Error())
		case models.ErrMissingRequiredParameters:
			//400
			b.ReplyErrContent(w, r, http.StatusBadRequest, err.Error())
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
		Status: "success",
	})
}

// GetAll list all
func (b *Building) GetAll(w http.ResponseWriter, r *http.Request) {
	data := &models.BuildingGetParams{}
	//check
	rows, err := data.GetAll(b.Storage)
	//chk
	if err != nil {
		//404
		b.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		return
	}
	//good
	render.JSON(w, r, Response{
		Status: "success",
		Result: rows,
		Total:  len(rows),
	})
}

// GetOne get 1 row per id
func (b *Building) GetOne(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingGetOne(strings.TrimSpace(chi.URLParam(r, "id")))
	//chk
	if data.ID == "" {
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	//check
	row, err := data.Get(b.Storage)
	//chk
	if err != nil {
		//404
		b.ReplyErrContent(w, r, http.StatusNotFound, err.Error())
		return
	}
	//good
	render.JSON(w, r, Response{
		Status: "success",
		Result: row,
	})
}

// Delete remove from store
func (b *Building) Delete(w http.ResponseWriter, r *http.Request) {
	data := models.NewBuildingDelete(strings.TrimSpace(chi.URLParam(r, "id")))
	//chk
	if data.ID == "" {
		//400
		b.ReplyErrContent(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	//chk
	if err := data.Delete(b.Storage); err != nil {
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
		Status: "success",
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
	info := struct {
		Application string `json:"application"`
		BuildTime   string `json:"buildTime"`
		Commit      string `json:"commit"`
		Release     string `json:"release"`
		Now         string `json:"now"`
	}{
		configs.Application,
		configs.BuildTime,
		configs.Commit,
		configs.Release,
		time.Now().Format(time.RFC3339),
	}
	render.JSON(w, r, info)
}
