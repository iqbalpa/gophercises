package utils

import (
	"strings"
)

func ArgParser(args []string) map[string]string {
	res := make(map[string]string)

	for _, arg := range args {
		arg = strings.TrimPrefix(arg, "-")
		kv := strings.Split(arg, "=")
		res[kv[0]] = kv[1]
	}

	return res
}
