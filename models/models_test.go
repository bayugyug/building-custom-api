package models_test

import (
	"context"
	"fmt"

	"github.com/bayugyug/building-custom-api/driver"
	"github.com/bayugyug/building-custom-api/models"
	"github.com/bayugyug/building-custom-api/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST Building API Service", func() {

	//init
	var ctx context.Context
	var store *driver.Storage

	BeforeEach(func() {
		ctx = context.Background()
		store = driver.NewStorage()
	})

	Context("Valid parameters", func() {

		Context("Create record", func() {
			It("should return ok", func() {
				params := &models.BuildingCreateParams{
					Name:    "test building name",
					Address: "address of the building name",
					Floors:  []string{"floor-a", "floor-b", "floor-c"},
				}
				pid, err := params.Create(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data ok")

			})
		})

		Context("Update record", func() {
			It("should return ok", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands",
					Address: "Marina Boulevard",
					Floors:  utils.Helper{}.SeedDataFloors(),
				}
				pid, err := params.Create(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before update ok")

				uparams := &models.BuildingUpdateParams{
					ID: pid,
					BuildingCreateParams: models.BuildingCreateParams{
						Name:    "marina-bay-sands",
						Address: "Marina Boulevard",
						Floors:  utils.Helper{}.SeedDataFloors(),
					},
				}

				err = uparams.Update(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				By("Update data ok")
			})
		})

		Context("Delete record", func() {
			It("should return ok", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands-1",
					Address: "Marina Boulevard-1",
					Floors:  utils.Helper{}.SeedDataFloors(),
				}
				pid, err := params.Create(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before delete ok")

				uparams := models.NewBuildingDelete(pid)
				err = uparams.Remove(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				By("Delete data ok")
			})
		})

		Context("Get 1 record", func() {
			It("should return ok", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands-2",
					Address: "Marina Boulevard-2",
					Floors:  utils.Helper{}.SeedDataFloors(),
				}
				pid, err := params.Create(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before get 1 row ok")

				uparams := models.NewBuildingGetOne(pid)
				row, err := uparams.Get(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(row.ID)).Should(BeNumerically(">", 0))
				By("Get 1 data ok")
			})
		})

		Context("Get list of records", func() {
			It("should return ok", func() {

				for i := 1; i <= 5; i++ {
					params := &models.BuildingCreateParams{
						Name:    fmt.Sprintf("marina-bay-sands-%05d", i),
						Address: fmt.Sprintf("Marina Boulevard-%05d", i),
						Floors:  utils.Helper{}.SeedDataFloors(),
					}
					pid, err := params.Create(ctx, store)
					if err != nil {
						Fail(err.Error())
					}
					Expect(len(pid)).Should(BeNumerically(">", 0))
					By("Create data before get list ok")
				}
				uparams := &models.BuildingGetOneParams{}
				rows, err := uparams.GetAll(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(rows)).Should(BeNumerically(">", 0))
				By("Get more data ok")
			})
		})

	}) // valid

	Context("Invalid parameters", func() {

		Context("Get list of records", func() {
			It("should error", func() {
				uparams := &models.BuildingGetOneParams{}
				rows, err := uparams.GetAll(ctx, store)
				Expect(err).To(HaveOccurred())
				Expect(len(rows)).To(Equal(0))
				By("Get more data empty as expected")
			})
		})

		Context("Create record", func() {
			It("should error", func() {
				params := &models.BuildingCreateParams{
					Address: "address of the building name",
					Floors:  []string{"floor-a", "floor-b", "floor-c"},
				}
				_, err := params.Create(ctx, store)
				Expect(err).To(HaveOccurred())
				By("Create data empty as expected")
			})
		})

		Context("Update record", func() {
			It("should error", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands",
					Address: "Marina Boulevard",
					Floors:  utils.Helper{}.SeedDataFloors(),
				}
				pid, err := params.Create(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before update ok")

				uparams := &models.BuildingUpdateParams{
					ID: pid + "::invalid",
					BuildingCreateParams: models.BuildingCreateParams{
						Name:    "marina-bay-sands",
						Address: "Marina Boulevard",
						Floors:  utils.Helper{}.SeedDataFloors(),
					},
				}

				err = uparams.Update(ctx, store)
				Expect(err).To(HaveOccurred())
				By("Update data empty as expected")
			})
		})

		Context("Delete record", func() {
			It("should error", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands-1a",
					Address: "Marina Boulevard-1a",
					Floors:  utils.Helper{}.SeedDataFloors(),
				}
				pid, err := params.Create(ctx, store)
				if err != nil {
					Fail(err.Error())
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before delete ok")

				uparams := models.NewBuildingDelete(pid + "no-id")
				err = uparams.Remove(ctx, store)
				Expect(err).To(HaveOccurred())
				By("Delete data empty as expected")
			})
		})

	}) // invalid
})
