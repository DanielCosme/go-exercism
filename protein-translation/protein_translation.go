package protein

import (
	"errors"
)

var ErrStop = errors.New("Err")
var ErrInvalidBase = errors.New("ERR")

func FromCodon(in string) (string, error) {
	switch res := protein(in); res {
	case "":
		return "", ErrInvalidBase
	case "STOP":
		return "", ErrStop
	default:
		return res, nil
	}
}

func FromRNA(in string) ([]string, error) {
	res := []string{}

	for i := 3; i-3 < len(in) && i <= 9; i += 3 {
		codon := in[i-3 : i]
		polypeptide, err := FromCodon(codon)
		if err != nil {
			if errors.Is(err, ErrStop) {
				err = nil
			}
			return res, err
		}

		res = append(res, polypeptide)
	}

	return res, nil
}

func protein(in string) string {
	switch in {
	case "AUG":
		return "Methionine"
	case "UUU", "UUC":
		return "Phenylalanine"
	case "UUA", "UUG":
		return "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine"
	case "UAU", "UAC":
		return "Tyrosine"
	case "UGU", "UGC":
		return "Cysteine"
	case "UGG":
		return "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "STOP"
	default:
		return ""
	}
}
