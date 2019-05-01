package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/bayugyug/building-custom-api/api/handler"
	"github.com/bayugyug/building-custom-api/api/routes"
	"github.com/bayugyug/building-custom-api/tools"

	"github.com/go-chi/chi"
	"github.com/icrowley/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST Building API Service", func() {

	//init
	var formdata string
	var router *chi.Mux
	routes.Service, _ = routes.NewAPIService(
		routes.WithSvcOptAddress(":8989"),
	)

	BeforeEach(func() {
		router = chi.NewRouter()
		router.Post("/v1/api/building", routes.Service.API.BuildCreate)
		router.Put("/v1/api/building", routes.Service.API.BuildingUpdate)
		router.Get("/v1/api/building", routes.Service.API.BuildingGet)
		router.Get("/v1/api/building/{id}", routes.Service.API.BuildingGetOne)
		router.Delete("/v1/api/building/{id}", routes.Service.API.BuildingDelete)
	})

	Context("Valid parameters", func() {

		Context("Create record", func() {
			It("should return ok", func() {
				formdata = tools.Seeder{}.Create()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Create data ok")
			})
		})

		Context("Update record", func() {
			It("should return ok", func() {
				buildingName := fmt.Sprintf("building-%s", fake.DigitsN(5))
				formdata = tools.Seeder{}.CreateWithName(buildingName)
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before update data ok")
				//update it
				pid, _ := response.Result.(string)
				formdata = tools.Seeder{}.Update(pid, buildingName)
				requestBody = bytes.NewReader([]byte(formdata))
				w2, body2 := testReq(router, "PUT", "/v1/api/building", requestBody)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusOK))
				Expect(response2.Code).To(Equal(http.StatusOK))
				By("Update data ok")
			})
		})

		Context("Get records", func() {
			It("should return ok", func() {

				//5 row
				for i := 0; i < 5; i++ {
					formdata = tools.Seeder{}.Create()
					requestBody := bytes.NewReader([]byte(formdata))
					w, body := testReq(router, "POST", "/v1/api/building", requestBody)
					var response handler.Response
					if err := json.Unmarshal(body, &response); err != nil {
						Fail(err.Error())
					}
					Expect(w.Code).To(Equal(http.StatusCreated))
					Expect(response.Code).To(Equal(http.StatusCreated))
					By("Add 1x1 data ok")
				}

				//get it
				w2, body2 := testReq(router, "GET", "/v1/api/building", nil)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusOK))
				Expect(response2.Code).To(Equal(http.StatusOK))
				By("Get more data ok")
			})
		})

		Context("Get 1 record", func() {
			It("should return ok", func() {
				formdata = tools.Seeder{}.Create()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add record before get 1 data ok")

				pid, _ := response.Result.(string)
				w2, body2 := testReq(router, "GET", "/v1/api/building/"+pid, nil)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusOK))
				Expect(response2.Code).To(Equal(http.StatusOK))
				By("Get 1 data ok")
			})
		})

		Context("Delete a record", func() {
			It("should return ok", func() {
				formdata = tools.Seeder{}.Create()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before remove data ok")

				pid, _ := response.Result.(string)
				w2, body2 := testReq(router, "DELETE", "/v1/api/building/"+pid, nil)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusOK))
				Expect(response2.Code).To(Equal(http.StatusOK))
				By("Remove data ok")
			})
		})
	}) // valid params

	Context("Invalid parameters", func() {

		Context("Create record", func() {
			It("should not create", func() {
				formdata = tools.Seeder{}.CreateWithEmptyName()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusPartialContent))
				Expect(response.Code).To(Equal(http.StatusPartialContent))
				By("Create data not done")
			})
		})

		Context("Create duplicate record", func() {
			It("should create again", func() {
				buildingName := fmt.Sprintf("building-%s", fake.DigitsN(5))
				formdata = tools.Seeder{}.CreateWithName(buildingName)
				w, body := testReq(router, "POST", "/v1/api/building", bytes.NewReader([]byte(formdata)))
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before duplicate ok")
				//do it again
				w2, body2 := testReq(router, "POST", "/v1/api/building", bytes.NewReader([]byte(formdata)))
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusConflict))
				Expect(response2.Code).To(Equal(http.StatusConflict))
				By("Duplicate data not allowed")
			})
		})

		Context("Create with missing name parameter", func() {
			It("should not create", func() {
				formdata = tools.Seeder{}.CreateWithEmptyName()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusPartialContent))
				Expect(response.Code).To(Equal(http.StatusPartialContent))
				By("Create data not done")
			})
		})

		Context("Try Get 1 record, no id parameter", func() {
			It("should not return data", func() {
				w, body := testReq(router, "GET", "/v1/api/building/no-id", nil)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(response.Code).To(Equal(http.StatusNotFound))
				By("Get data not found")
			})
		})

		Context("Update record with different ID", func() {
			It("should return not exists", func() {
				buildingName := fmt.Sprintf("building-%s", fake.DigitsN(5))
				formdata = tools.Seeder{}.CreateWithName(buildingName)
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before update data ok")
				//update it
				pid, _ := response.Result.(string)
				formdata = tools.Seeder{}.Update(pid+"-not-exists", buildingName)
				requestBody = bytes.NewReader([]byte(formdata))
				w2, body2 := testReq(router, "PUT", "/v1/api/building", requestBody)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusNoContent))
				Expect(response2.Code).To(Equal(http.StatusNoContent))
				By("Update data did not continue")
			})
		})

		Context("Update record with different name", func() {
			It("should return not exists", func() {
				buildingName := fmt.Sprintf("building-%s", fake.DigitsN(5))
				formdata = tools.Seeder{}.CreateWithName(buildingName)
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before update data ok")
				//update it
				pid, _ := response.Result.(string)
				formdata = tools.Seeder{}.Update(pid, buildingName+"-diff-name")
				requestBody = bytes.NewReader([]byte(formdata))
				w2, body2 := testReq(router, "PUT", "/v1/api/building", requestBody)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusConflict))
				Expect(response2.Code).To(Equal(http.StatusConflict))
				By("Update data did not continue")
			})
		})

		Context("Update record with missing required parameter", func() {
			It("should not update", func() {
				buildingName := fmt.Sprintf("building-%s", fake.DigitsN(5))
				formdata = tools.Seeder{}.CreateWithName(buildingName)
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before update data ok")
				//update it
				pid, _ := response.Result.(string)
				formdata = tools.Seeder{}.Update(pid, "")
				requestBody = bytes.NewReader([]byte(formdata))
				w2, body2 := testReq(router, "PUT", "/v1/api/building", requestBody)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusPartialContent))
				Expect(response2.Code).To(Equal(http.StatusPartialContent))
				By("Update data did not continue")
			})
		})

		Context("Delete a not existsing record", func() {
			It("should return ok", func() {
				formdata = tools.Seeder{}.Create()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response handler.Response
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(response.Code).To(Equal(http.StatusCreated))
				By("Add before remove data ok")

				pid, _ := response.Result.(string)
				w2, body2 := testReq(router, "DELETE", "/v1/api/building/"+pid+"not-exists", nil)
				var response2 handler.Response
				if err := json.Unmarshal(body2, &response2); err != nil {
					Fail(err.Error())
				}
				Expect(w2.Code).To(Equal(http.StatusNotFound))
				Expect(response2.Code).To(Equal(http.StatusNotFound))
				By("Remove data did not continue")
			})
		})
	}) // invalid params
})

// testReq
func testReq(router *chi.Mux, method, path string, body io.Reader) (*httptest.ResponseRecorder, []byte) {

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		Fail(err.Error())
	}

	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	respBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		Fail(err.Error())
	}

	return w, respBody
}