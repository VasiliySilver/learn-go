package main

import (
	"fmt"
	"strconv"
	"time"
)

// slice vs map
// 1 - сохранять большое количество PaymentInfo
// 2 - нужно находить конкретный элемент в PaymentInfo

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
	m.s = append(m.s, info)
}

func (m *PaymentModuleWithSlice) FindInfo(ID int) PaymentInfo {
	for _, info := range m.s {
		if info.ID == ID {
			return info
		}
	}

	return PaymentInfo{}
}

func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

func (m *PaymentModuleWithMap) FindInfo(ID int) PaymentInfo {
	info, ok := m.m[ID]
	if !ok {
		return PaymentInfo{}
	} else {
		return info
	}
}

func main() {
	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}

	iterations := 10_000

	before := time.Now()

	for i := 0; i < iterations; i++ {
		info := PaymentInfo{
			ID:          1,
			Description: "Desc: " + strconv.Itoa(i),
		}

		pSlice.AddInfo(info)
	}

	fmt.Println("slice add", time.Since(before))

	before = time.Now()

	for i := 0; i < iterations; i++ {
		info := PaymentInfo{
			ID:          1,
			Description: "Desc: " + strconv.Itoa(i),
		}

		pMap.AddInfo(info)
	}

	fmt.Println("map add", time.Since(before))

	// --------------------------------------
	before = time.Now()

	for i := 0; i < iterations; i++ {
		pSlice.FindInfo(i)
	}

	fmt.Println("slice find", time.Since(before))

	before = time.Now()

	for i := 0; i < iterations; i++ {
		pMap.FindInfo(i)
	}

	fmt.Println("map find", time.Since(before))

	// info1 := PaymentInfo{
	// 	ID:          10,
	// 	Description: "Some Description",
	// 	Usd:         10,
	// 	Cancelled:   false,
	// }

	// pSlice.AddInfo(info1)
	// pMap.AddInfo(info1)

	// pp.Println("pSlice: ", pSlice)
	// pp.Println("pMap: ", pMap)

	// i1 := pSlice.FindInfo(10)
	// i2 := pMap.FindInfo(10)

	// fmt.Println("i1", i1)
	// fmt.Println("i2", i2)
}
