package tools

import "fmt"

func IntToHex(i int64) string {
	h := fmt.Sprintf("%x", i)
	return h
}
