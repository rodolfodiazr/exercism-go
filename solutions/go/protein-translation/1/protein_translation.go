package protein

import "errors"

var (
    ErrStop error = errors.New("stop")
    ErrInvalidBase error = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
    if len(rna) % 3 != 0 {
        return []string{}, ErrInvalidBase
    }

    output := []string{}
    for i := 0; i <= len(rna) - 3; i += 3 {
        protein, err := FromCodon(
            string(rna[i]) + string(rna[i + 1]) + string(rna[i + 2]),
        )
        
        if err != nil {
            if errors.Is(err, ErrStop) {
                return output, nil
            }
            return output, err
        }

        output = append(output, protein)
    }
    return output, nil
}

func FromCodon(codon string) (string, error) {
    proteins := map[string]string{ 
        "AUG": "Methionine", 
        "UUU": "Phenylalanine", "UUC": "Phenylalanine",
        "UUA": "Leucine", "UUG": "Leucine",
        "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
        "UAU": "Tyrosine", "UAC": "Tyrosine", 
        "UGU": "Cysteine", "UGC": "Cysteine",
        "UGG": "Tryptophan",
        "UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
	}
    
    protein, ok := proteins[codon]
    if !ok {
    	return "", ErrInvalidBase
    }

    if protein == "STOP" {
        return "", ErrStop
    }
    return protein, nil
}
