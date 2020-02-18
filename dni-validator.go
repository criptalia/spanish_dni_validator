package spanish_dni_validator

import (
	"strconv"
	"strings"
)

const NIE_TYPES = "XYZ"
const DNINIE_CHECK_TABLE = "TRWAGMYFPDXBNJZSQVHLCKE"

const NIF_TYPES_WITH_LETTER_CHECK = "PQSW"
const NIF_TYPES_WITH_NUMBER_CHECK = "ABEH"

const NIF_LETTER_CHECK_TABLE = "JABCDEFGHI"

const DNI_REGEX = "^(?P<number>[0-9]{8})(?P<check>[A-Z])$"
const NIE_REGEX = "^(?P<type>[" + NIE_TYPES + "])(?P<number>[0-9]{7})(?P<check>[A-Z])$"
const OTHER_PERSONAL_NIF_REGEX = "^(?P<type>[KLM])(?P<number>[0-9]{7})(?P<check>[0-9A-Z])$"
const CIF_REGEX = "^(?P<type>[ABCDEFGHJNPQRSUVW])(?P<number>[0-9]{7})(?P<check>[0-9A-Z])$"

/**
 * Validate Spanish NIFS
 * Input is not uppercased, or stripped of any incorrect characters
 */
func IsValid(nif string) bool {
	return IsValidDni(nif) || IsValidNie(nif) || IsValidCif(nif) || IsValidOtherPersonalNif(nif)
}

/**
 * Validate Spanish NIFS given to persons
 */
func IsValidPersonal(nif string) bool {
return IsValidDni(nif) || IsValidNie(nif) || IsValidOtherPersonalNif(nif);
}

/**
 * Validate Spanish NIFS given to non-personal entities (e.g. companies, public corporations, ngos...)
 */
func IsValidEntity(nif string) bool {
	return IsValidCif(nif)
}

/**
 * DNI validation is pretty straight forward.
 * Just mod 23 the 8 digit number and compare it to the check table
 */
func IsValidDni(dni string) bool {
	result, matches, _ := regexMatch(DNI_REGEX, dni)
	if !result {
		return false
	}
	if "00000000" == matches["number"] {
		return false
	}

	number, ok := matches["number"]
	if !ok || number == "" {
		return false
	}
	check, ok := matches["check"]
	if !ok || check == "" {
		return false
	}

	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return false
	}

	return []rune(DNINIE_CHECK_TABLE)[numberInt % 23] == []rune(check)[0]
}

/**
 * NIE validation is similar to the DNI.
 * The first letter needs an equivalent number before the mod operation
 */
func IsValidNie(nie string) bool {
	result, matches, _ := regexMatch(NIE_REGEX, nie)
	if !result {
		return false
	}

	regType, ok := matches["type"]
	if !ok || matches["type"] == "" {
		return false
	}
	number, ok := matches["number"]
	if !ok || number == "" {
		return false
	}
	check, ok := matches["check"]
	if !ok || check == "" {
		return false
	}

	nieTypeInt := strings.Index(NIE_TYPES, regType)
	if nieTypeInt < 0 {
		return false
	}
	nieType := strconv.Itoa(nieTypeInt)

	nieString := nieType + number
	nieInt, err := strconv.Atoi(nieString)
	if err != nil {
		return false
	}

	return []rune(DNINIE_CHECK_TABLE)[nieInt % 23] == []rune(check)[0]
}

/**
 * Other personal NIFS are meant for temporary residents that do not qualify for a NIE but nonetheless need a NIF
 *
 * See references
 *
 * @see https://es.wikipedia.org/wiki/N%C3%BAmero_de_identificaci%C3%B3n_fiscal
 * @see https://es.wikipedia.org/wiki/C%C3%B3digo_de_identificaci%C3%B3n_fiscal
 */
func IsValidOtherPersonalNif(nif string) bool {
	result, matches, _ := regexMatch(OTHER_PERSONAL_NIF_REGEX, nif)
	if !result {
		return false
	}

	return isValidNifCheck(nif, matches)
}

/**
 * CIFS are only meant for non-personal entities
 *
 * See references
 *
 * @see https://es.wikipedia.org/wiki/N%C3%BAmero_de_identificaci%C3%B3n_fiscal
 * @see https://es.wikipedia.org/wiki/C%C3%B3digo_de_identificaci%C3%B3n_fiscal
 */
func IsValidCif(cif string) bool {
	result, matches, _ := regexMatch(CIF_REGEX, cif)
	if !result {
		return false
	}

	return isValidNifCheck(cif, matches)
}

func isValidNifCheck(nif string, matches map[string]string) bool {
	regType, ok := matches["type"]
	if !ok || matches["type"] == "" {
		return false
	}
	number, ok := matches["number"]
	if !ok {
		return false
	}
	check, ok := matches["check"]
	if !ok || check == "" {
		return false
	}

	splitStr := []rune(number)
	var split []int
	for _, val := range splitStr {
		valInt, err := strconv.Atoi(string(val))
		if err != nil {
			return false
		}
		split = append(split, valInt)
	}

	var even []int
	for key, val := range split {
		if key & 1 != 0 {
			even = append(even, val)
		}
	}
	var sumEven = 0
	for _, val := range even {
		sumEven += val
	}

	var odd []int
	for key, val := range split {
		if key & 1 == 0 {
			odd = append(odd, val)
		}
	}
	var sumOdd = 0
	for _, val := range odd {
		addVal := val * 2
		if addVal > 9 {
			addVal = addVal - 9
		}
		sumOdd += addVal
	}

	calculatedCheckDigit := (10 - (sumEven + sumOdd) % 10) % 10

	//Nifs with only letters
	if nifHasLetterCheck(regType, nif) {
		return checkNifLetter(calculatedCheckDigit, []rune(check)[0])
	}

	//Nifs with only numbers
	if nifHasNumberCheck(regType) {
		return checkNifNumber(calculatedCheckDigit, []rune(check)[0])
	}

	//Nifs that accept both
	return checkNifLetter(calculatedCheckDigit, []rune(check)[0]) || checkNifNumber(calculatedCheckDigit, []rune(check)[0])
}

func nifHasLetterCheck(nifType, nif string) bool {
	return -1 != strings.Index(NIF_TYPES_WITH_LETTER_CHECK, nifType) ||
		('0' == []rune(nif)[0] && '0' == []rune(nif)[1])
}

func checkNifLetter(calculatedCheckDigit int, checkDigit rune) bool {
	return []rune(NIF_LETTER_CHECK_TABLE)[calculatedCheckDigit] == checkDigit
}

func nifHasNumberCheck(nifType string) bool {
	return -1 != strings.Index(NIF_TYPES_WITH_NUMBER_CHECK, nifType)
}

func checkNifNumber(calculatedCheckDigit int, checkDigit rune) bool {
	return []rune(strconv.Itoa(calculatedCheckDigit))[0] == checkDigit
}
