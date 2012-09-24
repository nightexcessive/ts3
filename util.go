package ts3

import (
	"strings"
)

func escapeTS3String(str string) (returnString string) {
	returnString = str

	returnString = strings.Replace(returnString, "\\", "\\\\", -1)
	returnString = strings.Replace(returnString, "/", "\\/", -1)
	returnString = strings.Replace(returnString, " ", "\\s", -1)
	returnString = strings.Replace(returnString, "|", "\\p", -1)
	returnString = strings.Replace(returnString, "\a", "\\a", -1)
	returnString = strings.Replace(returnString, "\b", "\\b", -1)
	returnString = strings.Replace(returnString, "\f", "\\f", -1)
	returnString = strings.Replace(returnString, "\n", "\\n", -1)
	returnString = strings.Replace(returnString, "\r", "\\r", -1)
	returnString = strings.Replace(returnString, "\t", "\\t", -1)
	returnString = strings.Replace(returnString, "\v", "\\v", -1)

	return
}
