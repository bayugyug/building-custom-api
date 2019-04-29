package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/bayugyug/rest-building/controllers"
	"github.com/bayugyug/rest-building/utils"
	"github.com/go-chi/chi"
	"github.com/icrowley/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST Building API Service", func() {

	//init
	var formdata string
	var router *chi.Mux
	controllers.APIInstance, _ = controllers.NewAPIService(
		controllers.WithSvcOptAddress(":8989"),
	)

	BeforeEach(func() {
		router = chi.NewRouter()
		router.Post("/v1/api/building", controllers.APIInstance.API.BuildCreate)
		router.Put("/v1/api/building", controllers.APIInstance.API.BuildingUpdate)
		router.Get("/v1/api/building", controllers.APIInstance.API.BuildingGet)
		router.Get("/v1/api/building/{id}", controllers.APIInstance.API.BuildingGetOne)
		router.Delete("/v1/api/building/{id}", controllers.APIInstance.API.BuildingDelete)
	})

	Context("Create record", func() {
		It("should return ok", func() {
			formdata = utils.Helper{}.SeedData()
			requestBody := bytes.NewReader([]byte(formdata))
			req, err := http.NewRequest("POST", "/v1/api/building", requestBody)
			if err != nil {
				Fail(err.Error())
			}
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err := ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}
			var response controllers.APIResponse
			if err := json.Unmarshal(body, &response); err != nil {
				Fail(err.Error())
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response.Code).To(Equal(http.StatusOK))

			By("Create data ok")
		})
	})

	Context("Update record", func() {
		It("should return ok", func() {
			formdata = fmt.Sprintf(`{ "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
				fake.DigitsN(5),
				fake.DigitsN(5),
				fake.DigitsN(5),
			)
			requestBody := bytes.NewReader([]byte(formdata))
			req, err := http.NewRequest("POST", "/v1/api/building", requestBody)
			if err != nil {
				Fail(err.Error())
			}
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err := ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}
			var response controllers.APIResponse
			if err := json.Unmarshal(body, &response); err != nil {
				Fail(err.Error())
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response.Code).To(Equal(http.StatusOK))

			//update it
			formdata = fmt.Sprintf(`{ "id": "%s", "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
				response.Result,
				fake.DigitsN(5),
				fake.DigitsN(5),
				fake.DigitsN(5),
			)

			requestBody = bytes.NewReader([]byte(formdata))
			req, err = http.NewRequest("PUT", "/v1/api/building", requestBody)
			if err != nil {
				Fail(err.Error())
			}
			w = httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err = ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}

			var response2 controllers.APIResponse
			if err := json.Unmarshal(body, &response2); err != nil {
				Fail(err.Error())
			}
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response2.Code).To(Equal(http.StatusOK))

			By("Update data ok")
		})
	})

	Context("Get records", func() {
		It("should return ok", func() {

			for i := 0; i < 5; i++ {
				formdata = fmt.Sprintf(`{ "name": "building::a3","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
					fake.DigitsN(5),
					fake.DigitsN(5),
					fake.DigitsN(5),
				)
				requestBody := bytes.NewReader([]byte(formdata))
				req, err := http.NewRequest("POST", "/v1/api/building", requestBody)
				if err != nil {
					Fail(err.Error())
				}
				w := httptest.NewRecorder()
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(w, req)

				body, err := ioutil.ReadAll(w.Body)
				if err != nil {
					Fail(err.Error())
				}
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusOK))

			}

			//get it
			req, err := http.NewRequest("GET", "/v1/api/building", nil)
			if err != nil {
				Fail(err.Error())
			}
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err := ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}

			var response2 controllers.APIResponse
			if err := json.Unmarshal(body, &response2); err != nil {
				Fail(err.Error())
			}
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response2.Code).To(Equal(http.StatusOK))

			By("Get more data ok")
		})
	})

	Context("Get 1 record", func() {
		It("should return ok", func() {
			formdata = fmt.Sprintf(`{ "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
				fake.DigitsN(5),
				fake.DigitsN(5),
				fake.DigitsN(5),
			)
			requestBody := bytes.NewReader([]byte(formdata))
			req, err := http.NewRequest("POST", "/v1/api/building", requestBody)
			if err != nil {
				Fail(err.Error())
			}
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err := ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}
			var response controllers.APIResponse
			if err := json.Unmarshal(body, &response); err != nil {
				Fail(err.Error())
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response.Code).To(Equal(http.StatusOK))

			pid, _ := response.Result.(string)

			req, err = http.NewRequest("GET", "/v1/api/building/"+pid, nil)
			if err != nil {
				Fail(err.Error())
			}
			w = httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err = ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}

			var response2 controllers.APIResponse
			if err := json.Unmarshal(body, &response2); err != nil {
				Fail(err.Error())
			}
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response2.Code).To(Equal(http.StatusOK))

			By("Get 1 data ok")
		})
	})

	Context("Delete a record", func() {
		It("should return ok", func() {
			formdata = fmt.Sprintf(`{ "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
				fake.DigitsN(5),
				fake.DigitsN(5),
				fake.DigitsN(5),
			)
			requestBody := bytes.NewReader([]byte(formdata))
			req, err := http.NewRequest("POST", "/v1/api/building", requestBody)
			if err != nil {
				Fail(err.Error())
			}
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err := ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}
			var response controllers.APIResponse
			if err := json.Unmarshal(body, &response); err != nil {
				Fail(err.Error())
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response.Code).To(Equal(http.StatusOK))

			pid, _ := response.Result.(string)

			req, err = http.NewRequest("DELETE", "/v1/api/building/"+pid, nil)
			if err != nil {
				Fail(err.Error())
			}
			w = httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			body, err = ioutil.ReadAll(w.Body)
			if err != nil {
				Fail(err.Error())
			}

			var response2 controllers.APIResponse
			if err := json.Unmarshal(body, &response2); err != nil {
				Fail(err.Error())
			}
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response2.Code).To(Equal(http.StatusOK))
			By("Remove data ok")
		})
	})

})
