package env_kit

import (
	"os"
	"sort"
	"strings"
)

// FindAllEnvByPrefix
// prefix with "" will find all
//
//	find out all env by prefix
//	if not found will return nil
func FindAllEnvByPrefix(prefix string) map[string]string {
	var out map[string]string
	for _, e := range os.Environ() {
		if prefix == "" || strings.Index(e, prefix) == 0 {
			envSplit := strings.Split(e, "=")
			if len(envSplit) > 1 {
				if out == nil {
					out = make(map[string]string)
				}
				out[envSplit[0]] = envSplit[1]
			} else if len(envSplit) > 0 {
				if out == nil {
					out = make(map[string]string)
				}
				out[envSplit[0]] = ""
			} else {
				continue
			}
		}
	}
	if len(out) > 0 {
		return out
	}
	return nil
}

// FindAllEnvByPrefix4Print
//
//	find out all env by prefix for print
//	if not found will return ""
//
//nolint:golint,unused
func FindAllEnvByPrefix4Print(prefix string) string {
	var sb strings.Builder
	for _, e := range os.Environ() {
		if strings.Index(e, prefix) == 0 {
			sb.WriteString(e)
			sb.WriteString("\n")
		}
	}
	if sb.Len() > 0 {
		return sb.String()
	}
	return ""
}

// FindAllEnv4PrintAsSort
//
//	if not found will return ""
func FindAllEnv4PrintAsSort() string {
	return FindAllEnv4PrintAsSortJust(0)
}
func FindAllEnv4PrintAsSortJust(keySpaceJust uint) string {
	return FindEnv4PrintAsPrefixSortJust("", keySpaceJust)
}

// FindEnv4PrintAsPrefixSortJust
// prefix with "" will find all
// keySpaceJust 0 is not left justifying
//
//	find out all env and short
//	if not found will return ""
func FindEnv4PrintAsPrefixSortJust(prefix string, keySpaceJust uint) string {
	envMap := FindAllEnvByPrefix(prefix)
	if len(envMap) == 0 {
		return ""
	}
	var keys []string
	for k := range envMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		val := envMap[k]
		sb.WriteString(lJust(k, keySpaceJust, ' '))
		sb.WriteString("=")
		sb.WriteString(val)
		sb.WriteString("\n")

	}
	return sb.String()
}

func lJust(word string, width uint, fill rune) string {
	if width == 0 || word == "" {
		return word
	}

	wantWidth := int(width)
	wordLen := len(word)
	if wordLen >= wantWidth {
		return word
	}
	var sb strings.Builder
	sb.WriteString(word)
	between := wantWidth - wordLen
	for i := 0; i < between; i++ {
		sb.WriteRune(fill)
	}
	return sb.String()
}

// FindAllEnv4Print
//
//	if not found will return ""
//
//nolint:golint,unused
func FindAllEnv4Print() string {
	var sb strings.Builder
	for _, e := range os.Environ() {
		sb.WriteString(e)
		sb.WriteString("\n")
	}
	if sb.Len() > 0 {
		return sb.String()
	}
	return ""
}
