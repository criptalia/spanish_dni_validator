package spanish_dni_validator

import (
	"regexp"
)

var regexCache map[string]*regexp.Regexp

func regexMatch(regex, str string) (bool, map[string]string, error) {
	if regexCache == nil {
		regexCache = make(map[string]*regexp.Regexp)
	}

	compiledRegex, ok := regexCache[regex]
	if !ok {
		res, err := regexp.Compile(regex)
		if err != nil {
			return false, nil, err
		}
		compiledRegex = res
		regexCache[regex] = res
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
