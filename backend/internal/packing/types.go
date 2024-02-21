package packing

// PackMix is a struct that holds the result of the order pack calculations
type PackMix struct {
	Packs      map[int]int
	Total      int
	Count      int
	ExtraItems int
}
