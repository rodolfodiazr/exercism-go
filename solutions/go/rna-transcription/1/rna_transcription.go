package strand

import "strings"

func ToRNA(dna string) string {
    m := map[rune]rune{
        'G': 'C',
        'C': 'G',
        'T': 'A',
        'A': 'U',
    }
  
    var builder strings.Builder
    builder.Grow(len(dna))
    
    for _, c := range dna {
        r, ok := m[c]
        if !ok {
            return ""
        }
        builder.WriteRune(r)
    }
  
    return builder.String()
}
