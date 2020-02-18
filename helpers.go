package spanish_dni_validator

import (
	"regexp"
)

func pregMatch(regex, str string) (bool, map[string]string, error) {
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

/*
func pregMatch(regex, str string) (bool, map[string]string, error) {
	compiledRegex, err := pcre.Compile(regex, 0)
	if err != nil {
		return false, nil, errors.New(err.String())
	}
	match := compiledRegex.MatcherString(str, 0)
	if !match.Matches() {
		log.Warnf("no matches for %s on %s", regex, str)
		return false, nil, nil
	}
	result := make(map[string]string)
	groupsNumber := match.Groups()
	for i := 1; i <= groupsNumber; i++ {
		groupName := match.GroupString(i)
		if groupName != "" && match.NamedPresent(groupName) {
			result[groupName] = match.NamedString(groupName)
		}
	}
	return true, result, nil
}
*/