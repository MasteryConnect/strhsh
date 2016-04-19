package strhsh

import (
	"crypto/rand"

	"github.com/apex/log"
)

const AlphaNumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const Alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const Numeric = "0123456789"

func RandStr(strSize int, randType int) string {

	types := []string{"", "alphanum", "number", "alpha"}
	if len(types) < randType {
		log.Warn("invalid Random String type given for strhsh.RandStr()")
		return ""
	}
	rt := types[randType]
	var strMap string
	switch rt {
	default:
		log.Warn("incorrect type for RandStr. Choose 1-3 [alphanumeric, numeric, alpha")
	case "alphanum":
		strMap = AlphaNumeric
	case "number":
		strMap = Numeric
	case "alpha":
		strMap = Alpha
	}

	bytes := make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = strMap[v%byte(len(strMap))]
	}
	return string(bytes)
}
