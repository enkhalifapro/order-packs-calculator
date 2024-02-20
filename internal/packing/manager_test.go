package packing

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator(t *testing.T) {

	testData := []struct {
		Items               int
		Packs               []int
		ExpectedTotal       int
		ExpectedCount       int
		ExpectedExtraItem   int
		ExpectedCombination map[int]int
	}{
		{
			Items:             263,
			Packs:             []int{53, 31, 21},
			ExpectedTotal:     263,
			ExpectedCount:     7,
			ExpectedExtraItem: 0,
		},
		{
			Items:             263,
			Packs:             []int{53, 31, 23},
			ExpectedTotal:     263,
			ExpectedCount:     9,
			ExpectedExtraItem: 0,
			ExpectedCombination: map[int]int{
				23: 2,
				31: 7,
			},
		},
		{
			Items:             12001,
			Packs:             []int{5000, 2000, 1000, 500, 250},
			ExpectedTotal:     12250,
			ExpectedCount:     4,
			ExpectedExtraItem: 249,
		},
		{
			Items:             17820,
			Packs:             []int{5000, 2000, 1000, 500, 250},
			ExpectedTotal:     18000,
			ExpectedCount:     7,
			ExpectedExtraItem: 180,
		},
		{
			Items:             19110,
			Packs:             []int{5000, 2000, 1000, 500, 250},
			ExpectedTotal:     19250,
			ExpectedCount:     6,
			ExpectedExtraItem: 140,
		},
	}

	srv := NewManager(&logrus.Entry{})
	// Act
	for _, d := range testData {
		res := srv.CalculatePacks(d.Items, d.Packs)
		assert.Equal(t, res.Total, d.ExpectedTotal)
		assert.Equal(t, res.Count, d.ExpectedCount)
		assert.Equal(t, res.ExtraItems, d.ExpectedExtraItem)
		for k, v := range d.ExpectedCombination {
			assert.Equal(t, res.Packs[k], v)
		}
	}
}
