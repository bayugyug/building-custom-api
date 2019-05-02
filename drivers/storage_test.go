package drivers_test

import (
	"fmt"
	"time"

	"github.com/bayugyug/building-custom-api/drivers"
	"github.com/bayugyug/building-custom-api/models"
	"github.com/bayugyug/building-custom-api/tools"
	"github.com/icrowley/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST Building API Service::STORAGE", func() {

	//init
	var store *drivers.Storage
	var building models.BuildingData

	BeforeEach(func() {
		store = drivers.NewStorage()
		building = models.BuildingData{}
	})

	Context("Storage sanity checking", func() {

		Context("Create record", func() {
			It("should return ok", func() {
				name := fmt.Sprintf("building::%s", fake.DigitsN(15))
				pid := building.HashKey(name)
				record := &models.BuildingData{
					ID:      pid,
					Name:    name,
					Address: fmt.Sprintf("address::%s", fake.DigitsN(15)),
					Floors:  tools.Seeder{}.CreateFloors(),
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(gid).To(Equal(record.ID))
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
					Address: fmt.Sprintf("address::%s", fake.DigitsN(15)),
					Floors:  tools.Seeder{}.CreateFloors(),
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(gid).To(Equal(record.ID))
				By("Create data before update ok")

				//update old
				record = &models.BuildingData{
					ID:       pid,
					Name:     name,
					Address:  fmt.Sprintf("updated::address::%s", fake.DigitsN(15)),
					Floors:   tools.Seeder{}.CreateFloors(),
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
					Address: fmt.Sprintf("address::%s", fake.DigitsN(15)),
					Floors:  tools.Seeder{}.CreateFloors(),
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(gid).To(Equal(record.ID))
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

		Context("Get 1 record", func() {
			It("should return ok", func() {
				name := fmt.Sprintf("building::%s", fake.DigitsN(15))
				pid := building.HashKey(name)
				record := &models.BuildingData{
					ID:      pid,
					Name:    name,
					Address: fmt.Sprintf("address::%s", fake.DigitsN(15)),
					Floors:  tools.Seeder{}.CreateFloors(),
					Created: time.Now().Format(time.RFC3339),
				}
				gid := store.Set(pid, record)
				if gid == "" {
					Fail("Invalid generated ID")
				}
				Expect(gid).To(Equal(record.ID))
				By("Create data ok before get")

				//get 1
				data, err := store.One(pid)
				if err != nil {
					Fail(err.Error())
				}
				rec, oks := data.(*models.BuildingData)
				if !oks {
					Fail("Data conversion failed")
				}
				Expect(rec.ID).To(Equal(record.ID))
				Expect(rec.Name).To(Equal(record.Name))
				By("Get 1 data ok")
			})
		})

		Context("Get list of records", func() {
			It("should return ok", func() {
				for i := 1; i <= 10; i++ {
					name := fmt.Sprintf("building::%s", fake.DigitsN(15))
					pid := building.HashKey(name)
					record := &models.BuildingData{
						ID:      pid,
						Name:    name,
						Address: fmt.Sprintf("address::%s", fake.DigitsN(15)),
						Floors:  tools.Seeder{}.CreateFloors(),
						Created: time.Now().Format(time.RFC3339),
					}
					gid := store.Set(pid, record)
					if gid == "" {
						Fail("Invalid generated ID")
					}
					Expect(gid).To(Equal(record.ID))
				}

				By("Create data ok before get all records")

				//get all
				data, err := store.All()
				if err != nil || len(data) <= 0 {
					Fail(err.Error())
				}
				var all []*models.BuildingData
				for _, vv := range data {
					if row, valid := vv.(*models.BuildingData); valid {
						all = append(all, row)
					}
				}

				Expect(len(all)).To(Equal(10))
				By("Get all data ok")
			})
		})

	})
})
