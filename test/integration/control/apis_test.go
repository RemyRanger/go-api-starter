package control_test

import (
	"encoding/json"
	"go-api-starter/internal/common/models"
	"go-api-starter/internal/common/server"
	"go-api-starter/internal/services/apis/ports"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("APIs handler IT", func() {
	var (
		apiCreatedID uuid.UUID
	)

	Context("POST /apis", func() {
		It("should insert a new item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"title": "My first API",
				"version": "v1.0.0",
				"base_url": "https://test.com",
				"gateway_mode": "STRICT",
				"status": "ENABLE",
				"description": "test test test !"
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/apis", strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusCreated))

			var body ports.Api
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Title).To(Equal("My first API"))
			apiCreatedID = *body.ID

			// Verify the item was inserted into the database
			model := &models.Api{}
			err = db.First(model, apiCreatedID).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.Title).To(Equal("My first API"))
		})

		It("should reject wrong json body", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"title": "My first API",
				"version": "v1.0.0",
				"base_url": "https://test.com",
				"gateway_mode": "STRICT",dsdsqdq
				"status": "ENABLE",
				"description": "test test test !"
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/apis", strings.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))

			var body server.ErrorResponse
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())

			Expect(body.HTTPStatusCode).To(Equal(500))
			Expect(body.ErrorMessage).To(Equal("Internal Server Error"))
			Expect(body.Details).To(BeEmpty())
		})

		It("should reject wrong body fields format", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"title": "My first API",
				"version": "v1.0.0",
				"base_url": "https://test.com",
				"gateway_mode": "STRICTED",
				"status": "FETCHED",
				"description": "test test test !"
			}`
			req := httptest.NewRequest(http.MethodPost, "/v1/apis", strings.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusBadRequest))

			var body server.ErrorResponse
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())

			Expect(body.HTTPStatusCode).To(Equal(400))
			Expect(body.ErrorMessage).To(Equal("Bad Request"))
			Expect(body.Details).To(HaveLen(2))
			Expect(body.Details[0].Field).To(Equal("GatewayMode"))
			Expect(body.Details[0].Msg).To(Equal("Field validation failed on the 'oneof' tag"))
			Expect(body.Details[1].Field).To(Equal("Status"))
			Expect(body.Details[1].Msg).To(Equal("Field validation failed on the 'oneof' tag"))
		})
	})

	Context("GET /apis/{id}", func() {
		It("should fetch api by id", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/apis/"+apiCreatedID.String(), nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body ports.Api
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Title).To(Equal("My first API"))
		})

		It("should reject not found", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/apis/"+uuid.NewString(), nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusNotFound))

			var body server.ErrorResponse
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())

			Expect(body.HTTPStatusCode).To(Equal(404))
			Expect(body.ErrorMessage).To(Equal("Not Found"))
			Expect(body.Details).To(BeEmpty())
		})

		It("should reject wrong uuid", func() {
			// Prepare the HTTP GET request
			req := httptest.NewRequest(http.MethodGet, "/v1/apis/notuuid", nil)
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Context("PUT /apis/{id}", func() {
		It("should update item into the database", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"title": "My Updated API",
				"version": "v1.0.0",
				"base_url": "https://test.com",
				"gateway_mode": "STRICT",
				"status": "ENABLE",
				"description": "test test test !"
			}`
			req := httptest.NewRequest(http.MethodPut, "/v1/apis/"+apiCreatedID.String(), strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusOK))

			var body ports.Api
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body.Title).To(Equal("My Updated API"))
			apiCreatedID = *body.ID

			// Verify the item was inserted into the database
			model := &models.Api{}
			err = db.First(model, apiCreatedID).Error
			Expect(err).ToNot(HaveOccurred())
			Expect(model.Title).To(Equal("My Updated API"))
		})

		It("should reject with not found", func() {
			// Prepare the HTTP POST request
			reqBody := `{
				"title": "My Updated API",
				"version": "v1.0.0",
				"base_url": "https://test.com",
				"gateway_mode": "STRICT",
				"status": "ENABLE",
				"description": "test test test !"
			}`
			req := httptest.NewRequest(http.MethodPut, "/v1/apis/"+uuid.NewString(), strings.NewReader(reqBody))
			rec := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(rec, req)

			// Verify the HTTP response code
			Expect(rec.Code).To(Equal(http.StatusNotFound))

			var body server.ErrorResponse
			err := json.NewDecoder(rec.Body).Decode(&body)
			Expect(err).ToNot(HaveOccurred())

			Expect(body.HTTPStatusCode).To(Equal(404))
			Expect(body.ErrorMessage).To(Equal("Not Found"))
			Expect(body.Details).To(BeEmpty())
		})
	})
})
