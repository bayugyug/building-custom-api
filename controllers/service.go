package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bayugyug/building-custom-api/config"
	"github.com/bayugyug/building-custom-api/driver"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const (
	svcOptionWithHandler = "svc-opts-handler"
	svcOptionWithAddress = "svc-opts-address"
	svcOptionWithStore   = "svc-opts-store"
)

// APIInstance exposed var object
var APIInstance *APIService

// APIService the svc map
type APIService struct {
	API     *APIHandler
	Router  *chi.Mux
	Address string
}

// WithSvcOptHandler opts for handler
func WithSvcOptHandler(r *APIHandler) *config.Option {
	return config.NewOption(svcOptionWithHandler, r)
}

// WithSvcOptAddress opts for port#
func WithSvcOptAddress(r string) *config.Option {
	return config.NewOption(svcOptionWithAddress, r)
}

// WithSvcOptRedisHost opts for db connector
func WithSvcOptRedisHost(r string) *config.Option {
	return config.NewOption(svcOptionWithStore, r)
}

// NewAPIService service new instance
func NewAPIService(opts ...*config.Option) (*APIService, error) {

	//default
	svc := &APIService{
		Address: ":8989",
		API: &APIHandler{
			Storage: driver.NewStorage(),
			Context: context.Background(),
		},
	}

	//add options if any
	for _, o := range opts {
		//chk opt-name
		switch o.Name() {
		case svcOptionWithAddress:
			if s, oks := o.Value().(string); oks && s != "" {
				svc.Address = s
			}
		}
	} //iterate all opts

	//set the actual router
	svc.Router = svc.MapRoute()

	//good :-)
	return svc, nil
}

//Run run the http server based on settings
func (svc *APIService) Run() {

	//gracious timing
	srv := &http.Server{
		Addr:         svc.Address,
		Handler:      svc.Router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	//async run
	go func() {
		log.Println("Listening on port", svc.Address)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
			os.Exit(0)
		}

	}()

	//watcher
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	log.Println("Shutting down service...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

//MapRoute route map all endpoints
func (svc *APIService) MapRoute() *chi.Mux {

	// Multiplexer
	router := chi.NewRouter()

	// Basic settings
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
	)

	// Basic gracious timing
	router.Use(middleware.Timeout(60 * time.Second))

	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	router.Use(cors.Handler)

	router.Get("/", svc.API.Welcome)

	/*
		@end-points

		GET    /v1/api/building/:id
		POST   /v1/api/building
		PUT    /v1/api/building
		DELETE /v1/api/building/:id

	*/

	//end-points-mapping
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api",
			func(API *APIHandler) *chi.Mux {
				sr := chi.NewRouter()
				sr.Post("/building", API.BuildCreate)
				sr.Put("/building", API.BuildingUpdate)
				sr.Get("/building", API.BuildingGet)
				sr.Get("/building/{id}", API.BuildingGetOne)
				sr.Delete("/building/{id}", API.BuildingDelete)
				return sr
			}(svc.API))
	})

	return router
}

//SetContextKeyVal version context
func (svc *APIService) SetContextKeyVal(k, v string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), k, v))
			next.ServeHTTP(w, r)
		})
	}
}
