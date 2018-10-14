package protein

import (
	"errors"
)

var protein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	ErrInvalidBase = errors.New("ErrInvalidBase")
	STOP           = errors.New("Stop")
)

func FromCodon(codon string) (string, error) {
	v, ok := protein[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if v == "STOP" {
		return "", STOP
	}
	return v, nil
}

func FromRNA(rna string) ([]string, error) {

	end, front := 0, 3

	var proteinArr []string
	for front <= len(rna) {

		codon, err := FromCodon(rna[end:front])
		if err != nil {
			if err == STOP {
				return proteinArr, nil
			}
			return proteinArr, err
		}

		proteinArr = append(proteinArr, codon)

		end = front
		front += 3
	}

	return proteinArr, nil
}
