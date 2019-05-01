package drivers_test

import (
	"fmt"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
	"github.com/bayugyug/building-custom-api/models"
	"github.com/icrowley/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST Building API Service::STORAGE", func() {

	//init
	var store *drivers.Storage
	var building *models.BuildingData

	BeforeEach(func() {
		store = drivers.NewStorage()
		building = models.NewBuildingData()
	})

	Context("Storage sanity checking", func() {

		Context("Create record", func() {
			It("should return ok", func() {
				name := fmt.Sprintf("building::%s", fake.DigitsN(15))
				pid := building.HashKey(name)
				record := &models.BuildingData{
					ID:      pid,
					Name:    name,
					Address: "address of the building name",
					Floors:  []string{"floor-a", "floor-b", "floor-c"},
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data ok")
			})
		})

		Context("Update record", func() {
			It("should return ok", func() {
				name := fmt.Sprintf("building::%s", fake.DigitsN(15))
				pid := building.HashKey(name)
				record := &models.BuildingData{
					ID:      pid,
					Name:    name,
					Address: "address of the building name",
					Floors:  []string{"floor-a", "floor-b", "floor-c"},
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before update ok")

				//update old
				record = &models.BuildingData{
					ID:       pid,
					Name:     name,
					Address:  "updated address of the building name",
					Floors:   []string{"floor-a", "floor-b", "floor-c"},
					Modified: time.Now().Format(time.RFC3339),
				}
				gid = store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				//db check
				row, oks := store.Exists(pid)
				if !oks {
					Fail("Data not found")
				}
				//convert db data
				vrow, ok := row.(*models.BuildingData)
				if !ok {
					Fail("Data conversion failed")
				}
				Expect(vrow.Address).To(Equal(record.Address))
				By("Update ok")
			})
		})

		Context("Delete record", func() {
			It("should return ok", func() {
				name := fmt.Sprintf("building::%s", fake.DigitsN(15))
				pid := building.HashKey(name)
				record := &models.BuildingData{
					ID:      pid,
					Name:    name,
					Address: "address of the building name",
					Floors:  []string{"floor-a", "floor-b", "floor-c"},
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(len(pid)).Should(BeNumerically(">", 0))
				By("Create data before delete ok")

				//delete old
				row, oks := store.Exists(pid)
				if !oks {
					Fail("Data not found")
				}
				//convert db data
				if _, ok := row.(*models.BuildingData); !ok {
					Fail("Data conversion failed")
				}
				err := store.Unset(pid)
				Expect(err).To(BeZero())
				By("Delete ok")
			})
		})

	}) // valid

	Context("Invalid parameters", func() {

	}) // invalid
})
