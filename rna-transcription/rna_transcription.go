package strand

func ToRNA(dna string) (res string) {
	complement := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for _, v := range dna {
		res += string(complement[v])
	}

	return res
}
