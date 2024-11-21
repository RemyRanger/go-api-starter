package core_test

import (
	"context"
	"go-api-starter/internal/common/models"
	"go-api-starter/internal/services/apis/core"
	"go-api-starter/internal/services/apis/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Service APIs Tests", func() {
	var (
		ctrl    *gomock.Controller
		service *core.Service

		modelID uuid.UUID
		model   models.Api
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		model = models.Api{
			Title: "test",
		}
	})

	Describe("UpdateApi", func() {
		Context("with record in database", func() {
			BeforeEach(func() {
				repository := mocks.NewMockRepository(ctrl)

				modelID = uuid.New()
				repository.EXPECT().GetApi(gomock.Any(), modelID).Return(nil, nil)
				model.ID = modelID
				repository.EXPECT().UpsertApi(gomock.Any(), &model).Return(&model, nil)

				service = core.New(repository)
			})

			It("OK - should make 2 calls to repository", func() {
				result, err := service.UpdateApi(context.Background(), modelID, &model)

				Expect(*result).To(Equal(model))
				Expect(err).To(BeNil())
			})
		})

		Context("with no record in database", func() {
			BeforeEach(func() {
				repository := mocks.NewMockRepository(ctrl)

				modelID = uuid.New()
				repository.EXPECT().GetApi(gomock.Any(), modelID).Return(nil, gorm.ErrRecordNotFound)

				service = core.New(repository)
			})

			It("NOK - should propagate not found err on GetApi with no upsert call", func() {
				_, err := service.UpdateApi(context.Background(), modelID, &model)

				Expect(err).To(Equal(gorm.ErrRecordNotFound))
			})
		})
	})
})
