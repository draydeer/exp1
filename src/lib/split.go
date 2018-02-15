package lib

import (
	"strings"
)

func SplitKeyFn(separator string) func (string string) []string {
	escapedSeparator := "\\" + separator

	return func(key string) []string {
		if strings.Index(key, escapedSeparator) == - 1 {
			return strings.Split(key, separator)
		}

		keys := strings.Split(strings.Replace(key, escapedSeparator, "\f0", - 1), separator)

		for i, s := range keys {
			keys[i] = strings.Replace(s, "\f0", separator, - 1)
		}

		return keys
	}
}

func SplitKeyNFn(separator string) func (string string, n int) []string {
	escapedSeparator := "\\" + separator

	return func(key string, n int) []string {
		if strings.Index(key, escapedSeparator) == - 1 {
			return strings.SplitN(key, separator, n)
		}

		keys := strings.SplitN(strings.Replace(key, escapedSeparator, "\f0", n), separator, n)

		for i, s := range keys {
			keys[i] = strings.Replace(s, "\f0", separator, - 1)
		}

		return keys
	}
}
