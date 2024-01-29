package ports_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	cdvErrors "gitlab.com/cdv-projects/go-apis/internal/common/errors"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
	"gitlab.com/cdv-projects/go-apis/internal/services/flows"
	"gitlab.com/cdv-projects/go-apis/internal/services/flows/mocks"
	"gitlab.com/cdv-projects/go-apis/internal/services/flows/ports"
	"gorm.io/gorm"
)

func TestFlowsService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flows Service Suite")
}

func doGet(t *testing.T, mux *chi.Mux, url string) *httptest.ResponseRecorder {
	response := testutil.NewRequest().Get(url).WithAcceptJson().GoWithHTTPHandler(t, mux)
	return response.Recorder
}

var _ = Describe("FlowsHandler", func() {
	var (
		t                     *testing.T
		router                *chi.Mux
		ctrl                  *gomock.Controller
		sqlRepository         flows.SQL
		service               flows.Service
		flow1                 *models.Flow
		flowStatistique1      *models.FlowsStatistics
		talendFlowsStatistics *models.TalendFlowsStatistics
		newFlow               ports.NewFlow
		err                   error
	)

	BeforeEach(func() {
		router = chi.NewRouter()
		ctrl = gomock.NewController(GinkgoT())

		flow1 = &models.Flow{
			UUID: uuid.New(),
		}

		newFlow = ports.NewFlow{
			Description: "my description",
			FlowOrigin:  "FLOWORIGIN",
			JobName:     "FLOWJOBNAME",
			Format:      ports.NewFlowFormatAPIJSON,
			Logo:        "logo.jpg",
			Name:        "Name",
			Status:      ports.NewFlowStatusCREATED,
			TypeOffers:  ports.NewFlowTypeOffersNEUF,
		}

		var jobsOk []string
		var jobsError []string
		nbFlowsTotal := int64(20)
		flowStatistique1 = &models.FlowsStatistics{
			NbAds:        20,
			NbFlows:      20,
			NbFlowsTotal: &nbFlowsTotal,
			JobsError:    jobsError,
			JobsOk:       jobsOk,
		}

		talendFlowsStatistics = &models.TalendFlowsStatistics{
			SystemPid: 1125,
		}
	})

	JustBeforeEach(func() {
		service = flows.NewService(sqlRepository)
		ports.HandlerFromMuxWithBaseURL(ports.NewHTTPHandler(service), router, "/bo/v1")
	})

	Describe("CreateFlow", func() {
		Context("with no records in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				repositoryMock.EXPECT().CreateFlow(gomock.Any(), gomock.Any()).Return(flow1, nil)
				sqlRepository = repositoryMock
			})

			It("returns flow created", func() {
				rr := testutil.NewRequest().Post("/bo/v1/flows").WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var result ports.Flow
				err = json.NewDecoder(rr.Body).Decode(&result)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusCreated))
				Expect(err).To(BeNil())
				Expect(*result.ID).To(BeEquivalentTo(flow1.UUID.String()))
			})
		})

		Context("error create 1 flow in database", func() {
			It("returns an invalid request error", func() {
				newFlow.Status = "WRONG VALUE !!!"
				rr := testutil.NewRequest().Post("/bo/v1/flows").WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusBadRequest))
				Expect(err).To(BeNil())
			})
		})

		Context("with err in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				repositoryMock.EXPECT().CreateFlow(gomock.Any(), gomock.Any()).Return(nil, errors.New("database error"))
				sqlRepository = repositoryMock
			})

			It("returns an error", func() {
				rr := testutil.NewRequest().Post("/bo/v1/flows").WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusInternalServerError))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusInternalServerError)))
			})
		})
	})

	Describe("FindFlowByID", func() {
		Context("with 1 record in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var ID string = flow1.UUID.String()
				repositoryMock.EXPECT().GetFlow(gomock.Any(), ID).Return(flow1, nil)
				sqlRepository = repositoryMock
			})

			It("fetch 1 Flow", func() {
				rr := doGet(t, router, "/bo/v1/flows/"+flow1.UUID.String())

				var result ports.Flow
				err = json.NewDecoder(rr.Body).Decode(&result)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusOK))
				Expect(err).To(BeNil())
				Expect(*result.ID).To(BeEquivalentTo(flow1.UUID.String()))
			})
		})

		Context("with err in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var ID string = flow1.UUID.String()
				repositoryMock.EXPECT().GetFlow(gomock.Any(), ID).Return(nil, errors.New("database error"))
				sqlRepository = repositoryMock
			})

			It("returns an error", func() {
				rr := doGet(t, router, "/bo/v1/flows/"+flow1.UUID.String())

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusInternalServerError))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusInternalServerError)))
			})
		})

		Context("with err Not Found in database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var ID string = flow1.UUID.String()
				repositoryMock.EXPECT().GetFlow(gomock.Any(), ID).Return(nil, gorm.ErrRecordNotFound)
				sqlRepository = repositoryMock
			})

			It("returns an not found error", func() {
				rr := doGet(t, router, "/bo/v1/flows/"+flow1.UUID.String())

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusNotFound))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusNotFound)))
			})
		})
	})

	Describe("FindFlows", func() {
		Context("with 1 record in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var count int64 = 1
				repositoryMock.EXPECT().FindFlows(gomock.Any(), gomock.Any()).Return([]*models.Flow{flow1}, count, nil)
				sqlRepository = repositoryMock
			})

			It("fetch 1 flow", func() {
				rr := doGet(t, router, "/bo/v1/flows?order_by=uuid&page_size=10&page_token=1&sort_dir_desc=true")

				var result ports.ResponseFetchFlows
				err = json.NewDecoder(rr.Body).Decode(&result)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusOK))
				Expect(err).To(BeNil())
				Expect(len(*result.Flows)).To(BeEquivalentTo(1))
			})
		})

		Context("with err in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var count int64 = 0
				repositoryMock.EXPECT().FindFlows(gomock.Any(), gomock.Any()).Return(nil, count, errors.New("database error"))
				sqlRepository = repositoryMock
			})

			It("returns an error", func() {
				rr := doGet(t, router, "/bo/v1/flows")

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusInternalServerError))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusInternalServerError)))
			})
		})
	})

	Describe("FindJobsByFlowJobName", func() {
		Context("with 1 record in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var code string = "UBIFLOW"
				var count int64 = 1
				repositoryMock.EXPECT().FindJobsByFlowJobName(gomock.Any(), code, gomock.Any()).Return([]*models.TalendFlowsStatistics{talendFlowsStatistics}, count, nil)
				sqlRepository = repositoryMock
			})

			It("fetch 1 job", func() {
				rr := doGet(t, router, "/bo/v1/flows/UBIFLOW/jobs?page_size=10&page_token=1")

				var result ports.FlowJobs
				err = json.NewDecoder(rr.Body).Decode(&result)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusOK))
				Expect(err).To(BeNil())
				Expect(*result.Count).To(BeEquivalentTo(1))
				Expect(len(*result.FlowJobs)).To(BeEquivalentTo(1))
			})
		})

		Context("with err in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var code string = "UBIFLOW"
				var count int64 = 1
				repositoryMock.EXPECT().FindJobsByFlowJobName(gomock.Any(), code, gomock.Any()).Return(nil, count, errors.New("database error"))
				sqlRepository = repositoryMock
			})

			It("returns an error", func() {
				rr := doGet(t, router, "/bo/v1/flows/UBIFLOW/jobs?page_size=10&page_token=1")

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusInternalServerError))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusInternalServerError)))
			})
		})
	})

	Describe("UpdateFlow", func() {
		Context("update 1 Flow in database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var ID string = flow1.UUID.String()
				repositoryMock.EXPECT().UpdateFlow(gomock.Any(), gomock.Any(), ID).Return(flow1, nil)
				sqlRepository = repositoryMock
			})

			It("returns 1 product", func() {
				rr := testutil.NewRequest().Put("/bo/v1/flows/"+flow1.UUID.String()).WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var result ports.Flow
				err = json.NewDecoder(rr.Body).Decode(&result)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusCreated))
				Expect(err).To(BeNil())
				Expect(*result.ID).To(BeEquivalentTo(flow1.UUID.String()))
			})
		})

		Context("error update 1 Project in database", func() {
			It("returns an invalid request error", func() {
				newFlow.Status = "WRONG VALUE !!!"
				rr := testutil.NewRequest().Put("/bo/v1/flows/"+flow1.UUID.String()).WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusBadRequest))
				Expect(err).To(BeNil())
			})
		})

		Context("with err in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var ID string = flow1.UUID.String()
				repositoryMock.EXPECT().UpdateFlow(gomock.Any(), gomock.Any(), ID).Return(nil, errors.New("database error"))
				sqlRepository = repositoryMock
			})

			It("returns an error", func() {
				rr := testutil.NewRequest().Put("/bo/v1/flows/"+flow1.UUID.String()).WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusInternalServerError))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusInternalServerError)))
			})
		})

		Context("with err not found in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				var ID string = flow1.UUID.String()
				repositoryMock.EXPECT().UpdateFlow(gomock.Any(), gomock.Any(), ID).Return(nil, gorm.ErrRecordNotFound)
				sqlRepository = repositoryMock
			})

			It("returns an not found error", func() {
				rr := testutil.NewRequest().Put("/bo/v1/flows/"+flow1.UUID.String()).WithJsonBody(newFlow).GoWithHTTPHandler(t, router).Recorder

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusNotFound))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusNotFound)))
			})
		})
	})

	Describe("FindFlowsStatistics", func() {
		Context("with 1 record in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				repositoryMock.EXPECT().FindFlowsStatistics(gomock.Any()).Return(flowStatistique1, nil)
				sqlRepository = repositoryMock
			})

			It("fetch 1 FlowsStatistics", func() {
				rr := doGet(t, router, "/bo/v1/flows/stats")

				var result ports.FlowsStatistics
				err = json.NewDecoder(rr.Body).Decode(&result)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusOK))
				Expect(err).To(BeNil())
				Expect(result.NbAds).To(BeEquivalentTo(20))
			})
		})

		Context("with err in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				repositoryMock.EXPECT().FindFlowsStatistics(gomock.Any()).Return(nil, errors.New("database error"))
				sqlRepository = repositoryMock
			})

			It("returns an error", func() {
				rr := doGet(t, router, "/bo/v1/flows/stats")

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusInternalServerError))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusInternalServerError)))
			})
		})

		Context("with err not found in the database", func() {
			BeforeEach(func() {
				repositoryMock := mocks.NewMockSQL(ctrl)
				repositoryMock.EXPECT().FindFlowsStatistics(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
				sqlRepository = repositoryMock
			})

			It("returns an not found error", func() {
				rr := doGet(t, router, "/bo/v1/flows/stats")

				var httpError cdvErrors.ErrResponse
				err = json.NewDecoder(rr.Body).Decode(&httpError)

				Expect(rr.Code).To(BeEquivalentTo(http.StatusNotFound))
				Expect(err).To(BeNil())
				Expect(httpError.StatusText).To(BeEquivalentTo(http.StatusText(http.StatusNotFound)))
			})
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
