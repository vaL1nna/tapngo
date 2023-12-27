package utils

import (
	"fmt"
	"sort"
	"strings"
)

func BuildQueryString(params map[string]string) string {
	keys := make([]string, 0, len(params))

	for k, v := range params {
		if v == "" {
			continue
		}
		keys = append(keys, fmt.Sprintf("%s=%s", k, v))
	}

	sort.Strings(keys)
	return strings.Join(keys, "&")
}
