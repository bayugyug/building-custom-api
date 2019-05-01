package models_test

import (
	"context"
	"fmt"

	"github.com/bayugyug/building-custom-api/drivers"
	"github.com/bayugyug/building-custom-api/models"
	"github.com/bayugyug/building-custom-api/tools"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST Building API Service", func() {

	//init
	var ctx context.Context
	var store *drivers.Storage

	BeforeEach(func() {
		ctx = context.Background()
		store = drivers.NewStorage()
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
					Floors:  tools.Seeder{}.CreateFloors(),
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
						Floors:  tools.Seeder{}.CreateFloors(),
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
					Floors:  tools.Seeder{}.CreateFloors(),
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
					Floors:  tools.Seeder{}.CreateFloors(),
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
						Floors:  tools.Seeder{}.CreateFloors(),
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

		Context("Get list of records with empty list", func() {
			It("should error", func() {
				uparams := &models.BuildingGetOneParams{}
				rows, err := uparams.GetAll(ctx, store)
				Expect(err).To(HaveOccurred())
				Expect(len(rows)).To(Equal(0))
				By("Get more data empty as expected")
			})
		})

		Context("Create record with missing parameter", func() {
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

		Context("Update record with missing parameter id", func() {
			It("should error", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands",
					Address: "Marina Boulevard",
					Floors:  tools.Seeder{}.CreateFloors(),
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
						Floors:  tools.Seeder{}.CreateFloors(),
					},
				}

				err = uparams.Update(ctx, store)
				Expect(err).To(HaveOccurred())
				By("Update data empty as expected")
			})
		})

		Context("Delete record with missing parameter id", func() {
			It("should error", func() {
				params := &models.BuildingCreateParams{
					Name:    "marina-bay-sands-1a",
					Address: "Marina Boulevard-1a",
					Floors:  tools.Seeder{}.CreateFloors(),
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

		Context("Get a record not exists", func() {
			It("should error", func() {
				uparams := &models.BuildingGetOneParams{ID: "not-exists-id"}
				row, err := uparams.Get(ctx, store)
				Expect(err).To(HaveOccurred())
				Expect(row).To(BeZero())
				By("Get data empty as expected")
			})
		})
	}) // invalid
})
