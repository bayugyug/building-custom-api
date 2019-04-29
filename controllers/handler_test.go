package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	Context("Valid parameters", func() {

		Context("Create record", func() {
			It("should return ok", func() {
				formdata = utils.Helper{}.SeedData()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
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
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusOK))
				By("Add before update data ok")

				//update it
				formdata = fmt.Sprintf(`{ "id": "%s", "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
					response.Result,
					fake.DigitsN(5),
					fake.DigitsN(5),
					fake.DigitsN(5),
				)
				requestBody = bytes.NewReader([]byte(formdata))
				w2, body2 := testReq(router, "PUT", "/v1/api/building", requestBody)
				var response2 controllers.APIResponse
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
					formdata = fmt.Sprintf(`{ "name": "building::a3","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
						fake.DigitsN(5),
						fake.DigitsN(5),
						fake.DigitsN(5),
					)
					requestBody := bytes.NewReader([]byte(formdata))
					w, body := testReq(router, "POST", "/v1/api/building", requestBody)
					var response controllers.APIResponse
					if err := json.Unmarshal(body, &response); err != nil {
						Fail(err.Error())
					}
					Expect(w.Code).To(Equal(http.StatusOK))
					Expect(response.Code).To(Equal(http.StatusOK))
					By("Add 1x1 data ok")
				}

				//get it
				w2, body2 := testReq(router, "GET", "/v1/api/building", nil)
				var response2 controllers.APIResponse
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
				formdata = fmt.Sprintf(`{ "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
					fake.DigitsN(5),
					fake.DigitsN(5),
					fake.DigitsN(5),
				)

				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusOK))
				By("Add record before get 1 data ok")

				pid, _ := response.Result.(string)
				w2, body2 := testReq(router, "GET", "/v1/api/building/"+pid, nil)
				var response2 controllers.APIResponse
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
				formdata = fmt.Sprintf(`{ "name": "building::a1","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
					fake.DigitsN(5),
					fake.DigitsN(5),
					fake.DigitsN(5),
				)
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusOK))
				By("Add before remove data ok")

				pid, _ := response.Result.(string)
				w2, body2 := testReq(router, "DELETE", "/v1/api/building/"+pid, nil)
				var response2 controllers.APIResponse
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

		Context("Try Create record", func() {
			It("should not create", func() {
				formdata = utils.Helper{}.SeedDataEmptyName()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusPartialContent))
				By("Create data not done")
			})
		})

		Context("Try Update record, no name, no id parameter", func() {
			It("should not update", func() {
				formdata = utils.Helper{}.SeedDataEmptyName()
				requestBody := bytes.NewReader([]byte(formdata))
				w, body := testReq(router, "POST", "/v1/api/building", requestBody)
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusPartialContent))
				By("Update data not done")
			})
		})

		Context("Try Get 1 record, no id parameter", func() {
			It("should not return data", func() {
				w, body := testReq(router, "GET", "/v1/api/building/no-id", nil)
				var response controllers.APIResponse
				if err := json.Unmarshal(body, &response); err != nil {
					Fail(err.Error())
				}
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(response.Code).To(Equal(http.StatusNotFound))
				By("Get data not found")
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
