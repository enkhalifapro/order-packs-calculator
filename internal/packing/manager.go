package packing

import (
	"github.com/sirupsen/logrus"
	"sort"
)

// Manager is a struct that holds the packing functionalities
type Manager struct {
	logger *logrus.Entry
}

func NewManager(logger *logrus.Entry) *Manager {
	return &Manager{
		logger: logger,
	}

}

// CalculatePacks is a function that calculates the best combination of packs for order
func (mgr *Manager) CalculatePacks(items int, packs []int) PackMix {
	if len(packs) == 0 {
		return PackMix{}
	}

	// sort packs desc
	sort.Slice(packs, func(i, j int) bool {
		return packs[i] > packs[j]
	})

	m := getLeastPacks(items, packs)

	options := getPackCombinationOptions(m, items, packs)

	bestOption := getBestPackCombination(options)

	return bestOption
}

func getLeastPacks(items int, packs []int) map[int]int {
	m := make(map[int]int)
	distributed := 0
	for i := 0; i < len(packs); i++ {
		rem := (items - distributed) % packs[i]
		v := (items - distributed) / packs[i]
		m[packs[i]] = v
		distributed += packs[i] * v
		if rem > 0 && rem < packs[len(packs)-1] {
			m[packs[len(packs)-1]] += 1
		}
	}
	return m
}

func getPackCombinationOptions(leastPacks map[int]int, items int, packs []int) []PackMix {
	res := make([]PackMix, 0)

	total, count, mCpy := getMapTotal(leastPacks)
	extraItems := total - items
	if extraItems >= 0 {
		res = append(res, PackMix{
			Packs:      mCpy,
			Total:      total,
			Count:      count,
			ExtraItems: extraItems,
		})
	}
	for i := 0; i < len(packs); i++ {
		k := packs[i]
		v := leastPacks[k]
		rem := 0

		for j := 0; j < v; j++ {
			reduced := v - (j + 1)
			leastPacks[k] = reduced

			subItems := k
			reFill := subItems + rem

			m2, r := fill(reFill, packs[i+1:])
			rem = r
			for k1, v1 := range m2 {
				leastPacks[k1] += v1
			}

			total, count, mCpy := getMapTotal(leastPacks)
			extraItems := total - items
			if extraItems >= 0 {
				res = append(res, PackMix{
					Packs:      mCpy,
					Total:      total,
					Count:      count,
					ExtraItems: extraItems,
				})
			}

			if i == (len(packs) - 1) {
				if rem <= packs[len(packs)-1] {
					leastPacks[packs[len(packs)-1]] += 1
					total, count, mCpy := getMapTotal(leastPacks)
					extraItems := total - items
					if extraItems >= 0 {
						res = append(res, PackMix{
							Packs:      mCpy,
							Total:      total,
							Count:      count,
							ExtraItems: extraItems,
						})
					}

				}
			}
		}
	}
	return res
}

func getBestPackCombination(packs []PackMix) PackMix {
	bestComb := packs[0]
	for i := 1; i < len(packs); i++ {
		if packs[i].ExtraItems < bestComb.ExtraItems {
			bestComb = packs[i]
			continue
		}
		if packs[i].ExtraItems == bestComb.ExtraItems {
			if packs[i].Count < bestComb.Count {
				bestComb = packs[i]
				continue
			}
		}
	}
	return bestComb
}

func fill(items int, packs []int) (map[int]int, int) {
	m := make(map[int]int)
	if len(packs) == 0 {
		return m, items
	}
	rem := 0

	for i := 0; i < len(packs); i++ {
		rem = items % packs[i]
		v := (items - rem) / packs[i]
		m[packs[i]] = v
		items -= v * packs[i]
	}
	return m, rem
}

func getMapTotal(m map[int]int) (int, int, map[int]int) {
	totalItems := 0
	totalPacks := 0
	newM := make(map[int]int)
	for k, v := range m {
		totalPacks += v
		totalItems += k * v
		newM[k] = v
	}
	return totalItems, totalPacks, newM
}
