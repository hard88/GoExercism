package strand

import "strings"

//var transcriptions = map[string]string{
//	"G": "C",
//	"C": "G",
//	"T": "A",
//	"A": "U",
//}

func ToRNA(dna string) string {
	//var rna string
	//
	//for _, s := range dna {
	//	rna = rna + transcriptions[string(s)]
	//}

	//return rna

	rna := strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
		).Replace(dna)
	return rna
}
