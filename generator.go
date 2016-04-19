package strhsh

import (
	"crypto/rand"
	"errors"
)

const AlphaNumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const Alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const Numeric = "0123456789"

func RandStr(strSize int, randType int) (string, error) {
	types := []string{"", "alphanum", "number", "alpha"}
	if len(types) < randType {
		err := errors.New("invalid Random String type given for strhsh.RandStr()")
		return "", err
	}
	rt := types[randType]
	var strMap string
	switch rt {
	default:
		strMap = AlphaNumeric
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
	return string(bytes), nil
}
