package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{}
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0

	for _, v := range d {
		if !strings.ContainsRune("ACGT", v) {
			return nil, errors.New("Invalid nucleotide")
		}

		h[v]++
	}

	return h, nil
}
