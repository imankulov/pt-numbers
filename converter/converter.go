package converter

import (
	"fmt"
	"math"
	"strings"
)

type fold struct {
	ord      int
	singular string
	plural   string
}

// Do converts integer value value to string `out`. May return empty string
// and error, if input value is out of range
func Do(value int) string {
	if value == 0 {
		return "zero"
	}

	prefix := ""
	if value < 0 {
		value = -value
		prefix = "menus "
	}

	folds := []fold{
		fold{12, " trilhão", " trilhões"},
		fold{9, " bilhão", " bilhões"},
		fold{6, " milhão", " milhões"},
		fold{3, " mil", " mil"},
		fold{0, "", ""},
	}

	out := make([]string, 0, len(folds))
	for _, f := range folds {
		foldValue := value % int(math.Pow(10, float64(f.ord+3))) / int(math.Pow(10, float64(f.ord)))
		if foldValue == 0 {
			continue
		} else if foldValue == 1 {
			out = append(out, fmt.Sprintf("um%s", f.singular))
		} else {
			foldStringValue := convert0to999(foldValue)
			out = append(out, fmt.Sprintf("%s%s", foldStringValue, f.singular))
		}
	}

	return fmt.Sprintf("%s%s", prefix, strings.Join(out, " e "))
}

func convert0to999(value int) string {
	hundreds := value / 100
	rest := value % 100

	// process special cases (there's only, for 100)
	if value == 100 {
		return "cem"
	}

	restValue := convert0to99(rest)
	if hundreds == 0 {
		return restValue
	}

	hundredsMap := map[int]string{
		1: "cento",
		2: "duzentos",
		3: "trezentos",
		4: "quatrocentos",
		5: "quinhentos",
		6: "seiscentos",
		7: "setecentos",
		8: "oitocentos",
		9: "novecentos",
	}

	hundredsValue := hundredsMap[hundreds]
	if rest == 0 {
		return hundredsValue
	}

	return fmt.Sprintf("%s e %s", hundredsValue, restValue)
}

func convert0to99(value int) string {
	units := value % 10
	decades := value / 10 % 10

	// process special cases first
	specialCases := map[int]string{
		0:  "zero",
		10: "dez",
		11: "onze",
		12: "doze",
		13: "treze",
		14: "quatorze",
		15: "quinze",
		16: "dezesseis",
		17: "dezessete",
		18: "dezoito",
		19: "dezenove",
	}

	if specialCaseValue, ok := specialCases[value]; ok {
		return specialCaseValue
	}

	unitsMap := map[int]string{
		0: "",
		1: "um",
		2: "dois",
		3: "três",
		4: "quatro",
		5: "cinco",
		6: "seis",
		7: "sete",
		8: "oito",
		9: "nove",
	}

	unitValue := unitsMap[units]
	if decades == 0 {
		return unitValue
	}

	decadesMap := map[int]string{
		2: "vinte",
		3: "trinta",
		4: "quarenta",
		5: "cinquenta",
		6: "sessenta",
		7: "setenta",
		8: "oitenta",
		9: "noventa",
	}
	decadeValue := decadesMap[decades]

	if units == 0 {
		return decadeValue
	}

	return fmt.Sprintf("%s e %s", decadeValue, unitValue)
}
