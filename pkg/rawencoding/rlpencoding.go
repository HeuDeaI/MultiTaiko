package rawencoding

import (
	"fmt"
	"strconv"
	"strings"
)

func StringRlpEcnode(value string) string {
	value = strings.TrimPrefix(value, "0x")

	if number, _ := strconv.ParseUint(value, 16, 64); number < 0x80 && number > 0x00 {
		return value
	}

	bytesLength := (len(value) + 1) / 2
	prefix := 0x80 + bytesLength

	stringPrefix := fmt.Sprintf("%x", prefix)

	return stringPrefix + value
}

func ListRlpEcnode(values []string) string {
	var sb strings.Builder
	for _, value := range values {
		sb.WriteString(StringRlpEcnode(value))
	}
	result := sb.String()

	bytesLength := (len(result) + 1) / 2
	prefix := 0xc0 + bytesLength

	stringPrefix := fmt.Sprintf("%x", prefix)

	return stringPrefix + result
}
