package strand

var dna2rna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {

	var rna []rune
	for _, w := range dna {
		rna = append(rna, dna2rna[w])
	}

	return string(rna)
}
