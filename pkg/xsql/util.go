package xsql

import (
	"strconv"
	"strings"
)

func IdsToStr(ids []int64) string {
	str := strings.Builder{}
	for i, id := range ids {
		if i > 0 {
			str.WriteString(",")
		}
		str.WriteString(strconv.FormatInt(id, 10))
	}
	return str.String()
}
