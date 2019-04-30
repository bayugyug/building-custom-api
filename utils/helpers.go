package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

type Helper struct {
}

//RemoveIntDuplicates de-duplicate int list
func (h Helper) RemoveIntDuplicates(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

//RemoveStrDuplicates de-duplicate string list
func (h Helper) RemoveStrDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

//FormatSliceToIntMap convert slice of int to map
func (h Helper) FormatSliceToIntMap(all []int) map[int]int {
	bmap := make(map[int]int)
	for _, bv := range all {
		bmap[bv] = bv
	}
	return bmap
}

//UUID random uuid
func (h Helper) UUID() string {
	return uuid.New().String() + `-` + time.Now().Format("20060102-150405")
}

//HashMD5
func (h Helper) HashMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SeedData geneate dummy build-data
func (h Helper) SeedData() string {
	return fmt.Sprintf(`{ "name": "building-%s","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
		fake.DigitsN(12),
		fake.DigitsN(5),
		fake.DigitsN(5),
		fake.DigitsN(5),
	)
}

// SeedDataEmpty empty column
func (h Helper) SeedDataEmptyName() string {
	return fmt.Sprintf(`{ "name": "","address": "address::%s","floors": ["floor-%s","floor-%s"] }`,
		fake.DigitsN(5),
		fake.DigitsN(5),
		fake.DigitsN(5),
	)
}

// SeedDataFloors generate random floor list
func (h Helper) SeedDataFloors() []string {
	var floors []string
	t := rand.Intn(50)
	for i := 1; i <= t; i++ {
		floors = append(floors, fmt.Sprintf("floor-%s", fake.DigitsN(15)))
	}
	return floors
}
