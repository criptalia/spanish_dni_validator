package spanish_dni_validator

import "testing"

func TestIsValid_Valid(t *testing.T) {
	testData := append(validPersonalNifs(), validEntityNifs()...)
	for _, value := range testData {
		if !IsValid(value) {
			t.Errorf("%s should pass IsValid", value)
		}
	}
}

func TestIsValid_Invalid(t *testing.T) {
	testData := append(invalidNifs(), invalidPersonalNifs()...)
	testData = append(testData, invalidEntityNifs()...)
	for _, value := range testData {
		if IsValid(value) {
			t.Errorf("%s should not pass IsValid", value)
		}
	}
}

func TestIsValidPersonal_Valid(t *testing.T) {
	for _, value := range validPersonalNifs() {
		if !IsValidPersonal(value) {
			t.Errorf("%s should pass IsValidPersonal", value)
		}
	}
}

func TestIsValidPersonal_Invalid(t *testing.T) {
	testData := append(invalidPersonalNifs(), validEntityNifs()...)
	for _, value := range testData {
		if IsValidPersonal(value) {
			t.Errorf("%s should not pass IsValidPersonal", value)
		}
	}
}

func TestIsValidDni_Valid(t *testing.T) {
	for _, value := range validDnis() {
		if !IsValidDni(value) {
			t.Errorf("%s should pass IsValidDni", value)
		}
	}
}

func TestIsValidDni_Invalid(t *testing.T) {
	for _, value := range invalidDnis() {
		if IsValidDni(value) {
			t.Errorf("%s should not pass IsValidDni", value)
		}
	}
}

func TestIsValidNie_Valid(t *testing.T) {
	for _, value := range validNies() {
		if !IsValidNie(value) {
			t.Errorf("%s should pass IsValidNie", value)
		}
	}
}

func TestIsValidNie_Invalid(t *testing.T) {
	for _, value := range invalidNies() {
		if IsValidNie(value) {
			t.Errorf("%s should not pass IsValidNie", value)
		}
	}
}

func TestIsValidOtherPersonalNif_Valid(t *testing.T) {
	for _, value := range validOtherNifs() {
		if !IsValidOtherPersonalNif(value) {
			t.Errorf("%s should pass IsValidOtherPersonalNif", value)
		}
	}
}

func TestIsValidOtherPersonalNif_Invalid(t *testing.T) {
	for _, value := range invalidOtherNifs() {
		if IsValidOtherPersonalNif(value) {
			t.Errorf("%s should not pass IsValidOtherPersonalNif", value)
		}
	}
}

func TestIsValidEntity_Valid(t *testing.T) {
	for _, value := range validEntityNifs() {
		if !IsValidEntity(value) {
			t.Errorf("%s should pass IsValidEntity", value)
		}
	}
}

func TestIsValidEntity_Invalid(t *testing.T) {
	testData := append(invalidEntityNifs(), validPersonalNifs()...)
	for _, value := range testData {
		if IsValidEntity(value) {
			t.Errorf("%s should not pass IsValidEntity", value)
		}
	}
}

func TestIsValidCif_Valid(t *testing.T) {
	for _, value := range validEntityNifs() {
		if !IsValidCif(value) {
			t.Errorf("%s should pass IsValidCif", value)
		}
	}
}

func TestIsValidCif_Invalid(t *testing.T) {
	testData := append(invalidEntityNifs(), validPersonalNifs()...)
	for _, value := range testData {
		if IsValidCif(value) {
			t.Errorf("%s should not pass IsValidCif", value)
		}
	}
}

func invalidNifs() []string {
	return []string{
		//Garbage
		"AAAAAAAAA",
		"999999999",
		"BBBBB",
		"1",
		"93471790C0",
		"00000000T",
	}
}

func validDnis() []string {
	return []string{
		//DNI
		"93471790C",
		"43596386R",
		"00000010X",
	}
}

func invalidDnis() []string {
	return []string{
		//DNI
		"93471790A",
		"43596386B",
		"00000010Y",
	}
}

func validNies() []string {
	return []string{
		//NIE
		"X5102754C",
		"Z8327649K",
		"Y4174455S",
	}
}

func invalidNies() []string {
	return []string{
		//NIE
		"X5102754A",
		"Z8327649B",
		"Y4174455C",
	}
}

func validOtherNifs() []string {
	return []string{
		//Other NIF
		"K9514336H",
	}
}

func invalidOtherNifs() []string {
	return []string{
		"M3118299M",
	}
}

func validPersonalNifs() []string {
	result := validDnis()
	result = append(result, validNies()...)
	return append(result, validOtherNifs()...)
}

func invalidPersonalNifs() []string {
	result := invalidDnis()
	result = append(result, invalidNies()...)
	return append(result, invalidOtherNifs()...)
}

func validEntityNifs() []string {
	return []string{
		//CIF
		"A58818501",
		"B65410011",
		"V7565938C",
		"V75659383",
		"F0605378I",
		"Q2238877A",
		"D40022956",
		"A05497920",
		"D80180961",
	}
}

func invalidEntityNifs() []string {
	return []string{
		//CIF
		"A5881850B",
		"B65410010",
		"V75659382",
		"V7565938B",
		"F06053787",
		"Q22388770",
		"D4002295J",
	}
}


