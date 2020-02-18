package spanish_dni_validator

import (
	"regexp"
)

func regexMatch(regex, str string) (bool, map[string]string, error) {
	compiledRegex, err := regexp.Compile(regex)
	if err != nil {
		return false, nil, err
	}
	match := compiledRegex.FindStringSubmatch(str)
	if match == nil {
		return false, nil, nil
	}
	result := make(map[string]string)
	for i, name := range compiledRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return true, result, nil
}
